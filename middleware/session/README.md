
# 自己实现一个 Session 的功能
## 先写接口
[Session](./practice/session.go) # Session 接口
```go
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}
```

[Manager](./practice/manager.go) # Session Manager(Session管理器) 接口
```go
type Manager interface {
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionID string) (session Session, err error)
}
```

## 写接口的实现