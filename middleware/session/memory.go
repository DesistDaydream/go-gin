package session

import (
	"fmt"
)

// Data is
// type Data struct {
// 	//
// 	Data   map[string]interface{}
// 	rwLock sync.RWMutex
// }

// Get session.Data 支持的操作,根据给定的 key 获取值
func (d *Data) Get(keys string) (values interface{}, err error) {
	// 获取读锁
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()

	value, ok := d.Data["key"]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}

	return value, nil
}

// Set session.Data 支持的操作,根据给定的 k/v 设定这些值
func (d *Data) Set(keys string, value interface{}) {
	// 获取读锁
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	d.Data["key"] = value
}

// Del session.Data 支持的操作,根据给定的 key，删除对应的 k/v 对
func (d *Data) Del(keys string, value interface{}) (err error) {
	// 获取读锁
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	delete(d.Data, "key")
	return
}
