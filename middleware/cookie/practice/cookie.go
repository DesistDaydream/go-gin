// 这是一个 cookie 练习，要满足以下目标
// 2个路由，/login 和 /home
// /login 用于设置 cookie，/home 是访问查看信息的请求
// 在请求 /home 之前，先运行中间件以检验是否存在 cookie
// 若不存在则访问 /home 错误，直到访问了 /login 之后，才可以正常访问 /home
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// 设置 cookie
	c.SetCookie("admin", "123456", 30, "/", "*", false, true)
	c.String(200, "登陆成功")
}

func home(c *gin.Context) {
	c.String(200, "欢迎回家")
}

func cookieMiddleware(c *gin.Context) {
	// 获取客户端是否携带 cookie
	cookie, err := c.Cookie("admin")
	fmt.Println(cookie)
	if err != nil {
		// c.Redirect(http.StatusMovedPermanently, "/login")
		c.String(http.StatusBadRequest, "%v 不存在", cookie)
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", login)

	r.Use(cookieMiddleware)
	{
		r.GET("/home", home)
	}

	r.Run(":8080")
}
