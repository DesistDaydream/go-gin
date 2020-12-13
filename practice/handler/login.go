package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginGet 登录界面 GET 请求处理
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "Hello Care Dailyer",
	})
}

// LoginPost 登录界面 POST 请求处理
func LoginPost(c *gin.Context) {
	fmt.Println("用户名为：", c.PostForm("username"))
	fmt.Println("密码为：", c.PostForm("password"))

	c.Redirect(http.StatusMovedPermanently, "/order")
}
