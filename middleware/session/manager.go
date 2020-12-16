package session

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	// SessionCookieName 是 SessionID 在 Cookie 中对应的 key
	SessionCookieName = "session_id"
	// SessionContextName 是 SessionData 在 gin 上下文中对应的 key
	SessionContextName = "session"
)

// Mgr 全局变量
var Mgr *Manager

// Data 表示一个用户的 SessionData 应该具有的属性
type Data struct {
	ID   string
	Data map[string]interface{}
	// 读写锁，锁的是上面的 Data
	rwLock sync.RWMutex
}

// NewData 实例化 Data
func NewData(id string) *Data {
	return &Data{
		ID:   id,
		Data: make(map[string]interface{}, 8),
	}
}

// Manager 是一个全局的 Session 管理
type Manager struct {
	Session map[string]*Data
	rwLock  sync.RWMutex
}

// InitManager 初始化 Manager
func InitManager() {
	Mgr = &Manager{
		Session: make(map[string]*Data, 1024),
	}
}

// GetSessionData 根据 SessionID 找到对应的 Data
func (m *Manager) GetSessionData(sessionID string) (d *Data, err error) {
	// 取之前加锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	//
	d, ok := m.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return
}

// CreateSession 创建一条 Session 记录
func (m *Manager) CreateSession() (d *Data) {
	// 造一个 SessionID
	uuidObj := uuid.NewV4()

	// 造一个和 sessionID 对应的 SessionData
	d = NewData(uuidObj.String())

	// 返回 SessionData
	return
}

// Middleware 实现一个 gin 框架的中间件，这里是一个中间件处理的逻辑
// 所有流经此中间件的请求，它的上下文中肯定会有一个 session
func Middleware(m *Manager) gin.HandlerFunc {
	if m == nil {
		panic("must call InitManager() before use it!")
	}
	return func(c *gin.Context) {
		fmt.Println("中间件开始认证")
		var d *Data
		// 从请求的 Cookie 中获取 SessionID
		sessionID, err := c.Cookie(SessionCookieName)
		if err != nil {
			// 无 SessionID 的话，给这个用户创建一个新的 SessionData，同时分配一个 SessionID
			d = m.CreateSession()
		}

		// 有 SessionID 的话，根据 SessionID 去 Session 的大仓库中获取对应的 SessionData
		d, err = m.GetSessionData(sessionID)
		if err != nil {
			// 根据用户传过来的 SessionID 在大仓库中取不到 SessionData
			d = m.CreateSession()
			// 更新用户 Cookie 中保存的那个 SessionID
			sessionID = d.ID
		}

		// 如何实现让后续所有的处理请求的方法都唔那个拿到 SessionData

		// 利用 gin 的 c.Set，然后中间件中 c.Next
		c.Set(SessionContextName, d)

		// 在 gin 框架中，要回写 Cookie 必须在处理请求的函数返回之前
		c.SetCookie(SessionCookieName, sessionID, 20, "/", "datalake.cn", false, true)
		c.Next()
	}
}
