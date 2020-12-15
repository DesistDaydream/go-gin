package session

// Session 定义一个 Session 服务应该具有的属性
type Session struct {
	SessionID   string
	SessionData map[string]interface{}
}
