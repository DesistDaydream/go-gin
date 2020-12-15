package session

import (
	"fmt"
	"sync"
)

// SessionData 表示一个具体的用户 Session 数据
type SessionData struct {
	ID   string
	Data map[string]interface{}
	// 读写锁，锁的是上面的 Data
	rwLock sync.RWMutex
}

// NewSessionData 实例化 SessionData
func NewSessionData(id string) *SessionData {
	return &SessionData{
		ID:   id,
		Data: make(map[string]interface{}, 8),
	}
}

// Manager 是一个全局的 Session 管理
type Manager struct {
	Session map[string]SessionData
	rwLock  sync.RWMutex
}

// GetSessionData 根据 SessionID 找到对应的 SessionData
func (m *Manager) GetSessionData(sessionID string) (sd SessionData, err error) {
	// 取之前加锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	//
	sd, ok := m.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return
}

// CreateSession 创建一条 Session 记录
func (m *Manager) CreateSession() {
	// 造一个 SessionID

	// 造一个和 sessionID 对应的 SessionData

	// 返回 SessionData

}
