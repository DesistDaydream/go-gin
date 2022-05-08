package handler

import (
	"fmt"
	"io"
	"time"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Message 是一条消息应该具有的基本属性
type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time string `json:"time"`
}

// NewMessage 实例化 Message
func NewMessage() *Message {
	return &Message{
		Name: "DesistDaydream",
		Body: "Hello World",
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}

// HandleJSON 首页界面处理
func HandleJSON(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.html", nil)
	fmt.Println("处理 JSON 数据")
	m := NewMessage()
	switch c.Request.Method {
	case "GET":
		// 将 struct 中的数据转换为 JSON 格式
		jsonData, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 响应 JSON 格式的默认值
		fmt.Fprintf(c.Writer, "%v\n", string(jsonData))
	case "POST":
		// 模拟下面这样的 curl 请求，程序将会根据 Request Body 中的内容替换 Message 结构体数据中的值，并返回结构体中的数据
		// 这就好比请求一个需要 TOKEN 的 API，我们只有使用正确的 TOKEN，才可以获取想要的信息
		// curl -XPOST http://172.19.42.248:8080/json -d '{"name":"lichenhao"}'
		//
		// 读取 Request 的 Body
		RequestBody, _ := io.ReadAll(c.Request.Body)
		fmt.Printf("请求体为：%v\n", string(RequestBody))
		// 将 Request Body 的 JSON 格式转换为 struct 类型，并将 struct 中的值替换为 JSON 中的值
		// 注意，struct 中仅传入一个 key 的值，则 struct 中也只有一个属性的值被替代，其他属性的值保持不变
		err := json.Unmarshal(RequestBody, m)
		if err != nil {
			fmt.Fprintf(c.Writer, "请检查 Body，格式不正确或数据类型不对")
			return
		}
		fmt.Printf("请求体转换为 struct 后的值为：%v\n", m)

		// 根据传入的 请求体 的值，判断认证是否成功
		// 比如现在假设，只有传入 {"name":"DesistDaydream"} 这个请求体时，才会响应结构体的数据给给客户端。
		switch m.Name {
		case "DesistDaydream":
			// 认证正确，将 struct 类型数据转换为 JSON 格式数据并响应给客户端
			jsonData, err := json.Marshal(m)
			if err != nil {
				fmt.Fprintf(c.Writer, "序列化出错，请使用其他数据格式的 Body")
			}
			fmt.Fprint(c.Writer, string(jsonData))
		default:
			fmt.Fprintf(c.Writer, "你好 %v,认证失败，请重试\n", m.Name)
		}
	default:
		fmt.Fprintf(c.Writer, "仅接受 GET 与 POST\n")
	}
}
