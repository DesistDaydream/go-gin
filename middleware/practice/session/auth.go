package main

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/GoGin/middleware/session"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 校验用户是否已登录的中间件
func AuthMiddleware(c *gin.Context) {
	fmt.Println("验证中间件开始验证是否有 Session")
	tmpD, _ := c.Get(session.SessionContextName)
	// 拿到 SessionData
	d := tmpD.(*session.Data)
	value, err := d.Get("isLogin")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	isLogin, ok := value.(bool)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	if !isLogin {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Next()
}
