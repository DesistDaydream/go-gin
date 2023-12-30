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
	c.SetCookie("admin", "123456", 30, "/", "", false, true)
	c.String(200, "Login Success!")
}

func home(c *gin.Context) {
	c.String(200, "欢迎回家")
}

// authCookieMiddleware 认证中间件，用于拦截到 /home 的请求，并检验 cookie 信息，成功才允许访问
func authCookieMiddleware(c *gin.Context) {
	// 检查客户端是否携带 cookie，如果有 cookie，则判断 cookie 的值是否满足要求
	if cookie, err := c.Cookie("admin"); err == nil {
		switch cookie {
		// 如果 cookie 的值为 123456，则返回主程序继续执行 Handler
		case "123456":
			return
		// 如果 cookie 的值不是 123456，则响应错误信息，且中断 Handler 的后续处理
		default:
			c.String(http.StatusUnauthorized, "cookie 值错误")
			c.Abort()
		}
	} else {
		// 如果客户端请求没有携带 cookie，则响应错误信息，且中断 Handler 的后续处理
		c.String(http.StatusUnauthorized, "未登录")
		// ！！！重要！！！
		// Abort() 用来防止调用待处理的 Handler。请注意，这不会停止当前的 Handler。
		// 假设您有一个授权功能的中间件，用于验证当前请求是否得到授权。如果授权失败(例如：密码不匹配)，则需要调用 Abort() 以确保不调用该请求的其余处理程序。
		// 用白话说就是：若验证失败，不再调用后续的 Handler 进行处理。否则就算验证失败，依然会返回 Handler 函数继续处理后续代码。因为这个中间件只是调用链的一部分，调用完成后，还是会回到 Handler 中继续处理后面的代码
		// 这样的话，就算验证失败，依然可以访问。但是我们想要的效果是失败之后，就不让访问了
		c.Abort()
	}

}

func main() {
	r := gin.Default()

	r.GET("/login", login)

	r.Use(authCookieMiddleware)
	{
		r.GET("/home", home)
	}

	r.Run(":8080")
}
