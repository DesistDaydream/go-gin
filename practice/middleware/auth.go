package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// MiddleWare 中间件逻辑
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件 auth 开始认证了")
		c.Set("请求", "中间件")
		// 执行中间件
		c.Next()

		// 中间件执行完后续的一些事
		status := c.Writer.Status()
		fmt.Println("中间件认证完毕", status)
		t2 := time.Since(t)
		fmt.Println("消耗时间", t2)
	}
}
