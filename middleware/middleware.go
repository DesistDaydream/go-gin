package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MiddleWare 中间件逻辑
func MiddleWare(c *gin.Context) {
	// 声明一个当前时间，以便计算各种代码逻辑的耗时
	t := time.Now()
	fmt.Println("中间件开始执行")

	// 设置变量到 gin.Context 结构体的 key 中，可以通过 Get() 取值
	c.Set("request", "这是一个中间件样例 XD")
	// 模拟中间件执行，让中间件执行2秒，然后统计执行时长
	time.Sleep(2 * time.Second)
	fmt.Println("中间件消耗时间：", time.Since(t))

	// c.Next() 是中间件处理中很重要的一环。该方法将会挂起中间件的执行程序，等待 Handler 完成之后，再执行 c.Next() 下面的代码。
	// 如果不写 c.Next() 那么会是这种输出结果：
	// 中间件开始执行
	// Handler 共消耗时间： 2.00028165s
	// 中间件+Handler 共消耗时间： 2.000404446s
	// 获取了中间件中写入到 gin.Context 中的值： 这是一个中间件样例 XD
	//
	// 如果写了 c.Next() ，那么会是这种输出结果：
	// 中间件开始执行
	// Handler 共消耗时间： 2.000796193s
	// 获取了中间件中写入到 gin.Context 中的值： 这是一个中间件样例 XD
	// 中间件+Handler 共消耗时间： 5.001351866s
	//
	// 所以如果将 c.Set() 写到 c.Next() 后面，此时已经挂起，正常的页面处理逻辑中就无法获取 c.Set() 设置的值。
	// 现在再去看 gin.Default()，其中注册的 Logger() 与 Recovery() 这俩中间件内部，也是调用了 c.Next()
	// 这个挂起逻辑主要是为了一些特殊需求所准备的，比如在 Handler 处理完客户端的请求后，再执行一些统计相关的操作，以便统计这次请求中的一些信息，比如本例中统计的各种执行时间。
	c.Next()

	// 中间件执行完后续的一些事,比如统计程序执行时间
	fmt.Println("中间件+Handler 共消耗时间：", time.Since(t))
}

// PartialMiddleWare 局部中间件逻辑
// 这个中间件声明了一个返回值，与上面的 MiddleWare() 只是写法不同，效果都是一样的，这里就是展现一下两种不同的声明和调用方法
func PartialMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("局部中间件开始执行")
	}
}

func main() {
	// 初始化 gin 引擎
	r := gin.Default()

	// 这是一个最基本的 Handler
	r.GET("/no_middleware", func(c *gin.Context) {
		c.String(http.StatusOK, "没有使用中间件")
	})

	// r.Use() 将参数中定义的中间件附加到 router 中(也就是注册中间件)，通过该方法附加的中间件将包含在每个 Handler 中
	// 后面的每个 Handler，都会先执行 MiddleWare() 这个中间件
	r.Use(MiddleWare)
	// 使用 {} 是为了代码规范，不写也可以。写了之后更清晰得可以看出来， GET /middleware 或 /partial_middleware 之前，需要经过中间件处理
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 从中间中使用的 c.Set() 方法获取其值，并将中间件中定义的值打印到服务端，并响应给客户端。
			req, _ := c.Get("request")
			fmt.Println("获取了中间件中写入到 gin.Context 中的值：", req)
			c.String(200, "%v", req)
			// 让 Handler 运行3秒，模拟一下处理时间以便观察中间件的处理逻辑
			time.Sleep(3 * time.Second)
		})

		// 还可以在 Handler 的路径参数后面添加调用中间件，以实现局部中间件功能。
		// 本次调用局部中间件之前，同样也会调用 MiddleWare()，只不过这次局部中间调用只在本 Handler 中生效
		r.GET("/partial_middleware", PartialMiddleWare(), func(c *gin.Context) {
			c.String(200, "这是一个局部的中间件样例")
		})
	}

	r.Run(":8080")
}
