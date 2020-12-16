package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler 登录页面处理器
func LoginHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "Hello Care Dailyer"})
	case "POST":
		//
		fmt.Println("用户名为：", c.PostForm("username"))
		fmt.Println("密码为：", c.PostForm("password"))
		// c.DefaultQuery()

		c.Redirect(http.StatusMovedPermanently, "/order")
	default:
		c.String(http.StatusNotFound, "本页面暂仅支持 GET 和 POST 请求\n")
	}
}
