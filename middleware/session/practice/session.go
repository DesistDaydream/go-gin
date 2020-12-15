package session

// Session 接口规范
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}

func session() {
	// 待研究 https://www.bilibili.com/video/BV1nE411y7oZ?p=10

}
