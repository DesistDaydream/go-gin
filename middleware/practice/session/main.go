// 这个练习，不使用现成的库，而是手动实现一个 Session 中间件

package main

import (
	"net/http"

	"github.com/DesistDaydream/GoGin/middleware/session"
	"github.com/DesistDaydream/GoGin/practice/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	// 待学习：https://www.bilibili.com/video/BV1B4411w7vv?p=142
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 使用中间件
	session.InitManager()
	r.Use(session.Middleware(session.Mgr))

	r.GET("/index", handler.IndexGet)
	r.Any("/login", handler.LoginHandler)
	// r.GET("/home", homeHandler)
	// r.GET("/vip", vipHandler)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run()
}
