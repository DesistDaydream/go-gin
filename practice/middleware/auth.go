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

	// 验证 url 参数中的 token 字段
	token, _ := c.GetQuery("token")

	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "请求未携带token，无权限访问",
			"data":   nil,
		})
		c.Abort()
		return
	}

	logrus.Info("当前请求的 Token 为：", token)

	// 执行中间件
	c.Next()

	// 中间件执行完后续的一些事
	logrus.Info("中间件+handler 共消耗时间：", time.Since(t))
}
