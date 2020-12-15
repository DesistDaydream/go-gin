package session

import (
	"fmt"
	// "github.com/DesistDaydream/GoGin/middleware/session"
)

// MemeorySession is
// type MemorySession struct {
// 	//
// }

// Get session.Data 支持的操作,根据给定的 key 获取值
func (s *Data) Get(keys string) (values interface{}, err error) {
	// 获取读锁
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	value, ok := s.Data["key"]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}

	return value, nil
}

// Set session.Data 支持的操作,根据给定的 k/v 设定这些值
func (s *Data) Set(keys string, value interface{}) {
	// 获取读锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.Data["key"] = value
}

// Del session.Data 支持的操作,根据给定的 key，删除对应的 k/v 对
func (s *Data) Del(keys string, value interface{}) (err error) {
	// 获取读锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.Data, "key")
	return
}
