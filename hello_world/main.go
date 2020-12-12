package main

import "github.com/gin-gonic/gin"

// *gin.Context结构体中包含http请求(*http.Request)信息、http响应(http.ResponseWriter)信息、等等
// 这个函数中输出的信息也可以输出到运行该程序的服务器当中，比如在其中用fmt.Print打印的信息会显示在服务器中
func PingGet(c *gin.Context) {
	// gin.H{} 接口，可以将其中的内容，传递给前端页面。具体方法详见 X.HTML() 方法
	h := gin.H{
		"message": "pong",
	}
	// X.JSON() 方法用来将 JSON 信息作为 http.ResponseWriter 响应给客户端。JSON内容由 h 来定义，如果想让内容为空，则把 h 替换为 nil 即可。
	// 同理，还有其他的比如 X.HTML() 方法，则是将指定的 html 文件作为 http.ResponseWriter 响应给客户端。
	c.JSON(200, h)
}

func main() {
	// 创建gin框架对象、配置gin默认中间件、返回一个gin框架对象
	r := gin.Default()

	// gin的 X.GET() 方法用来定义http的请求路由，还包括 X.POST、X.PUT 等等。第一个参数为uri路径，第二个参数为处理方式(i.e.当访问/ping页面时应该如何处理)
	// 这里将 *http.Request 处理方式从 main() 中分离，直接调用 PingGet 函数。
	r.GET("/ping", PingGet)

	// 运行程序，默认监听在 0.0.0.0:8080 上
	r.Run()
}
