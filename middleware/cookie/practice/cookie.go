// 这是一个 cookie 练习，要满足以下目标
// 2个路由，/login 和 /home
// /login 用于设置 cookie，/home 是访问查看信息的请求
// 在请求 /home 之前，先运行中间件以检验是否存在 cookie
// 若不存在则访问 /home 错误，直到访问了 /login 之后，才可以正常访问 /home
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// 设置 cookie
	c.SetCookie("admin", "123456", 30, "/", "*", false, true)
	c.String(200, "Login Success!")
}

func home(c *gin.Context) {
	c.String(200, "欢迎回家")
}

func authCookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端是否携带 cookie
		cookie, err := c.Cookie("admin")
		// fmt.Println(cookie)
		if err == nil {
			if cookie == "123456" {
				c.Next()
				return
			}
		}

		// 返回错误信息
		c.String(http.StatusUnauthorized, "未登录")
		// 防止调用待处理的处理程序。请注意，这不会停止当前的 Handler。
		// 假设您有一个授权中间件，用于验证当前请求是否得到授权。如果授权失败（例如：密码不匹配），请调用Abort以确保不调用该请求的其余处理程序。
		// 用白话说就是：若验证不通过，不再调用后续的 Handler 进行处理。
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", login)

	r.Use(authCookieMiddleware())
	{
		r.GET("/home", home)
	}

	r.Run(":8080")
}
