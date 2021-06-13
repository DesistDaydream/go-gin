package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleWare 中间件逻辑
func AuthMiddleWare(c *gin.Context) {
	t := time.Now()
	logrus.Info("认证中间件开始工作")

	// 执行中间件
	c.Next()

	// 中间件执行完后续的一些事
	logrus.Info("中间件+handler 共消耗时间", time.Since(t))
}
