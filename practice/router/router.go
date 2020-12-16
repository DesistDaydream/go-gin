package router

import (
	"github.com/DesistDaydream/GoGin/practice/handler"
	"github.com/DesistDaydream/GoGin/practice/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由，设定路由信息
func InitRouter(r *gin.Engine) {

	r.GET("/", handler.IndexGet)
	r.Any("/login", handler.LoginHandler)

	// 为本程序注册中间件，以便后续页面都只有在认证之后才可以访问
	r.Use(middleware.AuthMiddleWare)
	// 使用 {} 是为了代码规范，不写也可以
	{
		r.GET("/order", handler.OrderGet)
		r.POST("/order", handler.OrderPost)
		r.GET("/stock-in", handler.StockInGet)
		r.POST("/stock-in", handler.StockInPost)
		r.GET("/stock-out", handler.StockOutGet)
		r.POST("/stock-out", handler.StockOutPost)
		r.GET("/query", handler.QueryGet)
		r.POST("/query", handler.QueryPost)
		r.GET("/inventory", handler.CommodityGet)
	}
}
