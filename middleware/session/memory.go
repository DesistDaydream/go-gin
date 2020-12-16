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
func (data *Data) Get(keys string) (values interface{}, err error) {
	// 获取读锁
	data.rwLock.RLock()
	defer data.rwLock.RUnlock()

	value, ok := data.Data["key"]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}

	return value, nil
}

// Set session.Data 支持的操作,根据给定的 k/v 设定这些值
func (data *Data) Set(keys string, value interface{}) {
	// 获取读锁
	data.rwLock.Lock()
	defer data.rwLock.Unlock()
	data.Data["key"] = value
}

// Del session.Data 支持的操作,根据给定的 key，删除对应的 k/v 对
func (data *Data) Del(keys string, value interface{}) (err error) {
	// 获取读锁
	data.rwLock.Lock()
	defer data.rwLock.Unlock()
	delete(data.Data, "key")
	return
}
