// 这个练习，不使用现成的库，而是手动实现一个 Session 中间件

package main

import (
	"net/http"

	"github.com/DesistDaydream/go-gin/middleware/session"
	"github.com/gin-gonic/gin"
)

func main() {
	// 待学习：https://www.bilibili.com/video/BV1B4411w7vv?p=142
	r := gin.Default()
	r.LoadHTMLGlob("web/*")

	// 使用中间件
	session.InitManager("redis", "127.0.0.1:6379", "redis", "0")
	r.Use(session.Middleware(session.ManagerObject))

	// 设置路由
	r.GET("/index", indexHandler)
	r.Any("/login", loginHandler)
	r.GET("/order", AuthMiddleware, orderHandler)
	r.GET("/vip", AuthMiddleware, vipHandler)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run()
}
