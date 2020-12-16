package main

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
		// 用来绑定用户登录时填写的用户名和密码
		var u struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}
		// 将用户登陆时填写的表单数据与 u 绑定起来，以便后续使用
		c.ShouldBind(&u)

		// 处理 POST 请求的逻辑
		fmt.Println("用户名为：", u.Username)
		fmt.Println("密码为：", u.Password)
		// 判断用户名和密码是否正确
		if u.Username == "zn" && u.Password == "zn" {
			c.Redirect(http.StatusFound, "/order")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{"err": "用户名或密码错误"})
		}

	}
}

func indexHandler(c *gin.Context) {
	//
}

func homeHandler(c *gin.Context) {
	//
}

func vipHandler(c *gin.Context) {
	//
}
