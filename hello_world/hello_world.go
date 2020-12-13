package main

import "github.com/gin-gonic/gin"

// PingGet 处理具体逻辑。这个函数中输出的信息也可以输出到运行该程序的服务器当中，比如在其中用fmt.Print打印的信息会显示在服务器中
// gin.Context 结构体中包含 http/net 基本库中 *http.Request 和 http.ResponseWriter。使用实现该结构体的方法，即可对 request 与 response 进行处理。
func PingGet(c *gin.Context) {
	// X.JSON() 方法用来将 JSON 信息作为 http.ResponseWriter 响应给客户端。
	// 第一个参数是响应给客户端的 http 响应码；第二个参数是响应给客户端的 body，是 JSON 格式的数据。
	// gin.H{} 就是 map[string]interface{} 的快捷方式，用来定义定义一些 JSON 格式的数据
	// 同理，还有其他的比如 X.HTML() 方法，则是将指定的 html 文件作为 http.ResponseWriter 响应给客户端。
	c.JSON(200, gin.H{"message": "pong"})

	// 当然还有一个更简单的方法，就是类似于 net/http 包的基本示例一样，直接传几个字符给 Response。
	// 这里面不会覆盖 c.JSON() ，而是会把内容附加到后面，与 c.JSON() 传给 Response 的内容一起响应给客户端
	c.String(200, "Hello")
}

func main() {
	// 可以使用 gin.Default() 或 gin.New() 创建 gin 引擎实例。区别在于 gin.Default() 使用 gin.New() 创建 engine 实例，并额外使用了 Logger() 和 Recover() 中间件。
	// Logger() 是负责进行打印并输出日志的中间件,方便开发者进行程序调试;比如当客户端访问 gin 开发的应用时，会输出本次访问的信息。
	// Recovery() 中间件的作如果程序执行过程中遇到panc中断了服务,则 Recovery会恢复程序执行,并返回服务器500内误。
	// 通常情况下,我们使用 gin.Defaul() 创建 Engine 实例。
	r := gin.Default()

	// gin 的 X.GET() 方法处理 GET 请求。一般称之为 handler(处理器)。还包括 X.POST、X.PUT 等等对应处理各种不同类型请求的 handler。
	// 第一个参数为请求的 uri 路径，第二个参数为处理方式(i.e.当访问/ping页面时应该如何处理)
	// 这里将 *http.Request 处理方式从 main() 中分离，直接调用 PingGet 函数。
	r.GET("/ping", PingGet)

	// 运行程序，默认监听在 0.0.0.0:8080 上
	r.Run("0.0.0.0:8080")
}
