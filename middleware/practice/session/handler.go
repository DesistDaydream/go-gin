package main

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/go-gin/middleware/session"
	"github.com/gin-gonic/gin"
)

// LoginHandler 登录页面处理器
func loginHandler(c *gin.Context) {
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
			tmpD, ok := c.Get(session.SessionContextName)
			if !ok {
				panic("session 中间件错误")
			}
			// 拿到 SessionData
			d := tmpD.(session.Data)
			d.Set("isLogin", true)
			d.Save()
			// 登录成功
			c.Redirect(http.StatusFound, "/order")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{"err": "用户名或密码错误"})
		}
	}
}

// IndexGet is
func indexHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
	case "POST":
	}
}

func orderHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "order.html", nil)
	case "POST":
	}
}

func vipHandler(c *gin.Context) {
	//
}
