package session

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// MemorySessionData 表示一个用户的 SessionData 应该具有的属性
type MemorySessionData struct {
	ID   string
	Data map[string]interface{}
	// 读写锁，锁的是上面的 Data
	rwLock sync.RWMutex
}

// NewMemorySessionData 实例化 RedisSessionData
func NewMemorySessionData(id string) Data {
	return &MemorySessionData{
		ID:   id,
		Data: make(map[string]interface{}, 8),
	}
}

// GetID is
func (m *MemorySessionData) GetID() string {
	return m.ID
}

// Get Data 支持的操作,根据给定的 key 获取值
func (m *MemorySessionData) Get(key string) (values interface{}, err error) {
	// 获取读锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	value, ok := m.Data[key]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}

	return value, nil
}

// Set Data 支持的操作,根据给定的 k/v 设定这些值
func (m *MemorySessionData) Set(key string, value interface{}) {
	// 获取读锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.Data[key] = value
}

// Del Data 支持的操作,根据给定的 key，删除对应的 k/v 对
func (m *MemorySessionData) Del(key string) {
	// 获取读锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.Data, key)
}

// Save 保存 SessionData
func (m *MemorySessionData) Save() {
}

// MemoryManager 是一个全局的 Session 管理
type MemoryManager struct {
	Session map[string]Data
	rwLock  sync.RWMutex
}

// NewMemoryManager 实例化存储 SessionData 的 RAM 后端
func NewMemoryManager() Manager {
	return &MemoryManager{
		Session: make(map[string]Data, 1024),
	}
}

// Init 初始化
func (m *MemoryManager) Init(addr string, options ...string) {
}

// GetSessionData 根据 SessionID 找到对应的 Data
func (m *MemoryManager) GetSessionData(sessionID string) (d Data, err error) {
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
func (m *MemoryManager) CreateSession() (d Data) {
	// 造一个 SessionID
	uuidObj := uuid.New()
	// 造一个和 sessionID 对应的 SessionData
	d = NewMemorySessionData(uuidObj.String())
	// 将创建的 SessionID 保存到 SessionData 中
	m.Session[d.GetID()] = d
	// 返回 SessionData
	return
}
