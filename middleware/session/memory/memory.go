package memory

import (
	"fmt"
	"sync"
)

// SessionData is
type SessionData struct {
	//
	Data   map[string]interface{}
	rwLock sync.RWMutex
}

// Get session.Data 支持的操作,根据给定的 key 获取值
func (sd *SessionData) Get(keys string) (values interface{}, err error) {
	// 获取读锁
	sd.rwLock.RLock()
	defer sd.rwLock.RUnlock()

	value, ok := sd.Data["key"]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}

	return value, nil
}

// Set session.Data 支持的操作,根据给定的 k/v 设定这些值
func (sd *SessionData) Set(keys string, value interface{}) {
	// 获取读锁
	sd.rwLock.Lock()
	defer sd.rwLock.Unlock()
	sd.Data["key"] = value
}

// Del session.Data 支持的操作,根据给定的 key，删除对应的 k/v 对
func (sd *SessionData) Del(keys string, value interface{}) (err error) {
	// 获取读锁
	sd.rwLock.Lock()
	defer sd.rwLock.Unlock()
	delete(sd.Data, "key")
	return
}
