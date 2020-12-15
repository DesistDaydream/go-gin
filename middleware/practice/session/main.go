// 这个练习，不使用现成的库，而是手动实现一个 Session 中间件

package main

import (
	"github.com/DesistDaydream/GoGin/middleware/session"
	"github.com/gin-gonic/gin"
)

func main() {
	// 待学习：https://www.bilibili.com/video/BV1B4411w7vv?p=142
	r := gin.Default()
	r.LoadHTMLFiles("templates/*")

	// 使用中间件
	session.InitManager()
	r.Use(session.Middleware(session.Mgr))

	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", vipHandler)
	r.Run()
}
