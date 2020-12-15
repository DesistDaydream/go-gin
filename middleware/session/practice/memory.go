package session

import (
	"sync"
)

// MemorySession is
type MemorySession struct {
	sessionID string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

// NewMemorySession 返回一个储存 Session 的内存引擎
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{}
	return s
}
