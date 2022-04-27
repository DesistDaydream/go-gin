package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleWare 中间件逻辑
func AuthMiddleWare(c *gin.Context) {
	t := time.Now()
	logrus.Info("认证中间件开始工作")

	// 从名为 token 的 Cookie 中获取值
	token, _ := c.Cookie("token")

	logrus.Info("当前请求的 Token 为：", token)

	// 验证 Token 是否存在
	if token == "" {
		logrus.Error("请求未携带token，无权限访问")
		c.HTML(http.StatusOK, "index.html", gin.H{"err": "请求未携带token，无权限访问"})
		c.Abort()
		return
	}

	// 验证 Token 内容
	_, err := ParseToken(token)
	if err != nil {
		logrus.Error("验证 Token 失败，原因：", err)
		c.HTML(http.StatusOK, "index.html", gin.H{"err": err})
		c.Abort()
		return
	}

	// 执行中间件
	c.Next()

	// 中间件执行完后续的一些事
	logrus.Info("中间件+handler 共消耗时间：", time.Since(t))
}
