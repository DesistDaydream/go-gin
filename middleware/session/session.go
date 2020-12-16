package session

import (
	"fmt"

	"github.com/DesistDaydream/GoGin/middleware/session/storage"
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

// Data is
type Data interface {
	GetID() string // 返回自己的SessionID
	Get(keys string) (values interface{}, err error)
	Set(keys string, value interface{})
	Del(keys string)
	Save()
}

// Manager 所有存储 SessionData 的后端类型都应该遵循的接口
type Manager interface {
	// 所有支持的后端都必须实现 Init()
	Init(addr string, options ...string)
	GetSessionData(string) (d Data, err error)
	CreateSession() (d Data)
}

// InitManager 初始化 Manager
// 根据 name 参数，选择使用什么类型的后端来存储 SessionData
func InitManager(name string, addr string, options ...string) {
	switch name {
	case "memory":
		ManagerObject = storage.NewMemoryManager()
	case "redis":
		ManagerObject = storage.NewRedisManager()
	}
	// 初始化 ManagerObject
	ManagerObject.Init(addr, options...)
}

// SessionMiddleware 实现一个 gin 框架的中间件，这里是一个中间件处理的逻辑
// 所有流经此中间件的请求，它的上下文中肯定会有一个 session
func SessionMiddleware(m Manager) gin.HandlerFunc {
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
