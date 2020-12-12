package main

import "github.com/gin-gonic/gin"

// PingGet 处理具体逻辑 *gin.Context 结构体中包含http请求(*http.Request)信息、http响应(http.ResponseWriter)信息、等等
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
	// 默认可以使用 gin.Default() 和 gin.New() 创建 gin 引擎实例。区别在于 gin.Default() 也适用 gin.New() 创建 engine 实例，但是会默认使用 Logger 和 Recover 中间件。
	// Logger 是负责进行打印并输出日志的中间件,方便开发者进行程序调试;比如当客户端访问 gin 开发的应用时，会输出本次访问的信息。
	// Recovery 中间件的作如果程序执行过程中遇到panc中断了服务,则 Recovery会恢复程序执行,并返回服务器500内误。
	// 通常情况下,我们使用默认的 gin.Defaul() 创建 Engine 实例。
	r := gin.Default()

	// gin的 X.GET() 方法用来定义http的请求路由，还包括 X.POST、X.PUT 等等。第一个参数为uri路径，第二个参数为处理方式(i.e.当访问/ping页面时应该如何处理)
	// 这里将 *http.Request 处理方式从 main() 中分离，直接调用 PingGet 函数。
	r.GET("/ping", PingGet)

	// 运行程序，默认监听在 0.0.0.0:8080 上
	r.Run("0.0.0.0:8088")
}
