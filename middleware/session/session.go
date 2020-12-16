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

// Data 本身是一个 struct，用来定义 SessionData 应该具有的属性
// 但是由于其要实现的不同类型的存储后端所存放 SessionData 的行为是不同的。Redis 还需要 Save 保存一下数据，所以也就需要一个 connection pool 属性；但是 内存的却不需要。既然属性不同，那么就无法公用一个 struct
// 所以不再约束 SessionData 中的属性，而是约束如何操作 SessionData。而 操作 SessionData 就是实现一些方法。
// 所有后端都应该可以对 SessionData 执行同样的操作，也就是遵循同一个接口
type Data interface {
	// 返回自己的SessionID
	GetID() string
	// 根据 SessionData 中的 key 获取对应的 value
	Get(keys string) (values interface{}, err error)
	// 设置 SessionData 中的 key/value对
	Set(keys string, value interface{})
	// 根据 SessionData 中的 key 删除对应的 key/value对
	Del(keys string)
	// 将 SessionData 中的数据持久化
	Save()
}

// Manager 本身是一个 struct，用来定义 Session管理器 应该具有的属性
// 但是由于其要实现的存储 SessionData 的后端是一个动态切换的，可以是 RAM、Redis、Mysql 等等
// 如果是结构体的话，其中的属性要适应所有类型的后端是不可能：
// 比如 RAM 不需要一个 connection pool，但是 Redis 和 mysql 至少需要先创建一个 connection pool(就是打开一个连接)；而 Redis 和 Mysql 的 connection pool 又是不同的~
// 所以不再约束这些后端应该具有的属性(也就是不再定义一个统一的的结构体),而是约束这些后端的行为(即定义一个统一的接口，这些后端都要实现这些方法，方法就是这些后端的行为)
// 这些后端的行为，就如下面所使，应该包括这三个：初始化、获取 SessionData、创建 SessionData。
// 所有存储 SessionData 的后端类型都应该遵循的接口
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
		// 用户的每次访问，都要重新设置以下 Cookie，主要是更新过期时间
		// 在 gin 框架中，要回写 Cookie 必须在处理请求的函数返回之前
		c.SetCookie(SessionCookieName, sessionID, 60, "/", "datalake.cn", false, true)
		c.Next()
	}
}
