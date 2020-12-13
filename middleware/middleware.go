package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MiddleWare 中间件逻辑
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("请求", "中间件")
		// 执行中间件。
		// 如果跟踪 gin.Default()，其中注册的 Logger() 与 Recovery() 这俩中间件内部，也是调用了 c.Next()
		c.Next()

		// 中间件执行完后续的一些事
		status := c.Writer.Status()
		fmt.Println("中间件执行", status)
		t2 := time.Since(t)
		fmt.Println("中间件消耗时间", t2)
	}
}

// PartialMiddleWare 局部中间件逻辑
func PartialMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("局部中间件开始执行")
		c.Set("请求", "中间件")
		// 执行中间件。
		// 如果跟踪 gin.Default()，其中注册的 Logger() 与 Recovery() 这俩中间件内部，也是调用了 c.Next()
		c.Next()

		// 中间件执行完后续的一些事
		status := c.Writer.Status()
		fmt.Println("局部中间件执行", status)
		t2 := time.Since(t)
		fmt.Println("局部中间件消耗时间", t2)
	}
}

func main() {
	// 初始化 gin 引擎
	r := gin.Default()

	r.GET("/no_middleware", func(c *gin.Context) {
		c.String(http.StatusOK, "No Use Middleware")
	})

	// 为本程序注册中间件，以便后续页面都只有在认证之后才可以访问
	r.Use(MiddleWare())
	// 使用 {} 是为了代码规范，不写也可以。写了之后更清晰得可以看出来， GET /order 之前，需要经过中间件处理
	{
		r.GET("/middleware", func(c *gin.Context) {
			c.String(http.StatusOK, "Middleware Example")
		})

		// 还可以在路由路径后面添加调用中间件，以实现局部中间件功能。
		// 本次调用局部中间件之前，同样也会调用 MiddleWare()
		r.GET("/partial_middleware", PartialMiddleWare(), func(c *gin.Context) {
			c.String(http.StatusOK, "Partial Middleware")
		})
	}

	// 运行 gin 程序
	r.Run(":8080")
}
