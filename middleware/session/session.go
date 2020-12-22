package session

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	// SessionCookieName 是 SessionID 在 Cookie 中对应的 key
	SessionCookieName = "session_id"
	// SessionContextName 是 SessionData 在 gin 上下文中对应的 key
	SessionContextName = "session"
)

// ManagerObject 全局变量
var ManagerObject Manager

// Data 本身是一个 struct，用来定义 SessionData 应该具有的属性。参见 MemorySessionData 这个结构体。
// MemorySessionData 这个结构体实现了 GetID()、Get()、Set()、Del() 这几个方法。但是现在又有一个新的事物，也要实现这几种方法，就是 RedisSessionData 这个结构体。
// Redis 与 Memory 存储 SessionData 的方式不同(也就是结构体里的属性不同)，但是对 SessionData 的操作都是那么几个方法。
// 所以这时候不再约束 XXXSessionData 中的属性要保持一致，而是约束如何操作 SessionData。而操作 SessionData 就是实现一些方法。
// 这时候，就需要抽象出来一个 Interface，包含两种方式都支持的操作。所有存储方式都应该遵循该接口的定义
// 此时如果有某些操作在某个事物上不支持，比如 Memory 不持支 Save，那么 MemorySessionData 实现了 Save 之后，直接返回就行，方法内不用写任何代码。
// 这里面的接口就很像在学习go时，练习接口那里，只不过这里接口包含的方法更多罢了。见：https://github.com/DesistDaydream/GoLearning/tree/master/practice/interface
type Data interface {
	// 返回自己的SessionID
	GetID() string
	// 根据 SessionData 中的 key 获取对应的 value
	Get(key string) (values interface{}, err error)
	// 设置 SessionData 中的 key/value对
	Set(key string, value interface{})
	// 根据 SessionData 中的 key 删除对应的 key/value对
	Del(key string)
	// 将 SessionData 中的数据持久化
	Save()
}

// Manager 本身是一个 struct，用来定义 SessionManager 应该具有的属性。参见 MemoryManager 这个结构体。
// 但是由于其要实现的存储 SessionData 的后端是一个动态切换的，可以是 RAM、Redis、Mysql 等等。
// 比如 RAM 不需要一个 connection pool，但是 Redis 和 mysql 至少需要先创建一个 connection pool(就是打开一个连接)；而 Redis 和 Mysql 的 connection pool 又是不同的~
// 所以不再约束这些后端应该具有的属性(也就是不再定义一个统一的的结构体),而是约束这些后端的行为(即定义一个统一的接口，这些后端都要实现这些方法，方法就是这些后端的行为)
// 这些后端的行为，就如下面所示，应该包括这三个：初始化、获取 SessionData、创建 SessionData。
// 所有存储 SessionData 的后端类型都应该遵循的接口
// 这里就涉及到一个扩展的思想，最一开始只有 MemoryManager，但是当需要扩展一个新的 Manager 时，就需要给这些 Manager 抽象成统一的接口。
// 只要这些不管是 Memory、Redis 还是其他的 都具有接口中定义的方法，那么其他人在调用接口的时候，也就不用再关注具体怎么实现的，只要能给我正确的返回值即可。
type Manager interface {
	// 用来初始化 connection pool。内存不用初始化，在方法中直接 return 即可，不用执行任何行为
	// 如果是 Redis，则初始化一个 Redis 的 connection pool。在初始化的时候，提供 Redis 的 IP、Port、密码等信息即可。
	// 所以，所谓的初始化，就是打开一个连接，Redis 就是连接上 Redis、MySQL 就是连接上 MySQL
	Init(addr string, options ...string)
	// 根据一个给定的 SessionID，拿到对应的 SessionData
	GetSessionData(string) (d Data, err error)
	// 如果一个 新用户 或者 Session 过期的用户访问，那么我们需要给他们创建一个 SessionID
	// 这个功能一般是放在一个中间件中，单独实现，这个中间件的作用就是拦截用户请求，并根据用户请求的参数，判断这个用户是否有合法的 SessionID
	// 如果没有的话，就为其创建或更新一下 SessionID，并分配 SessionData
	CreateSession() (d Data)
}

// InitManager 初始化 Manager
// 根据 name 参数，选择使用什么类型的后端来存储 SessionData
func InitManager(name string, addr string, options ...string) {
	switch name {
	case "memory":
		ManagerObject = NewMemoryManager()
	case "redis":
		ManagerObject = NewRedisManager()
	}
	// 初始化 ManagerObject
	ManagerObject.Init(addr, options...)
}

// Middleware 实现一个 gin 框架的中间件，这里是一个中间件处理的逻辑
// 所有流经此中间件的请求，它的上下文中肯定会有一个 session
func Middleware(m Manager) gin.HandlerFunc {
	if m == nil {
		panic("must call InitManager() before use it!")
	}
	return func(c *gin.Context) {
		fmt.Println("Session 处理中间件开始处理 Session")
		var d Data
		// 从请求的 Cookie 中获取 SessionID
		sessionID, err := c.Cookie(SessionCookieName)
		// 判断是否有 SessionID，根据有无进行不同的处理
		if err != nil {
			// 无 SessionID 的话，给这个用户创建一个新的 SessionData，同时分配一个 SessionID
			d = m.CreateSession()
			sessionID = d.GetID()
			fmt.Println("无 SessionID，创建一个 SessionData，并分配一个 SessionID", sessionID)
		} else {
			// 有 SessionID 的话，根据 SessionID 去 Session 的大仓库中获取对应的 SessionData
			d, err = m.GetSessionData(sessionID)
			if err != nil {
				// 根据用户传过来的 SessionID 在大仓库中取不到 SessionData。(比如 SessionID 过期或错误)
				d = m.CreateSession()
				// 更新用户 Cookie 中保存的那个 SessionID
				sessionID = d.GetID()
				fmt.Println("SessionID过期，分配一个新 ID", sessionID)
			}
			fmt.Println("SessionID 未过期", sessionID)
		}

		// 如何实现让后续所有的处理请求的方法都能拿到 SessionData？

		// 利用 gin 的 c.Set，然后中间件中 c.Next
		c.Set(SessionContextName, d)
		// 设置 Cookie。用户的每次请求都要重新设置以下 Cookie，主要是更新过期时间
		// 注意：在 gin 框架中，要回写 Cookie 必须在处理请求的函数返回之前，也就是在 c.Next() 之前
		c.SetCookie(SessionCookieName, sessionID, 60, "/", "datalake.cn", false, true)
		c.Next()
	}
}
