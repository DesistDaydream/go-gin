package session

// Manager Session管理器 接口规范
type Manager interface {
	// 初始化
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionID string) (session Session, err error)
}

func manager() {
	// 待研究 https://www.bilibili.com/video/BV1nE411y7oZ?p=10

}
