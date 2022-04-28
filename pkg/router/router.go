package router

import (
	"net/http"

	"github.com/DesistDaydream/go-gin/pkg/api"
	"github.com/DesistDaydream/go-gin/pkg/handler"
	"github.com/DesistDaydream/go-gin/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由，设定路由信息
func InitRouter(r *gin.Engine) {
	// 测试用
	r.Any("/header", handler.HandleHeader)
	r.Any("/json", handler.HandleJSON)

	// 设置 api v1 分组的路由，用于接受前端发起的 POST 请求
	v1 := r.Group("/api/v1")
	v1.POST("/login", api.LoginPost)

	v1.Use(middleware.AuthMiddleWare)
	{
	}

	// 设置静态资源路由，通常都是 GET 请求
	r.GET("/login", handler.LoginGet)

	// 注册 js 资源路由
	r.GET("/frontend/js/login.js", handler.JSGet)

	// 为本程序注册中间件，以便后续页面都只有在认证之后才可以访问
	r.Use(middleware.AuthMiddleWare)
	// 使用 {} 是为了代码规范，不写也可以
	{
		r.Any("/order", handler.OrderHandler)
		r.GET("/stock-in", handler.StockInGet)
		r.POST("/stock-in", handler.StockInPost)
		r.GET("/stock-out", handler.StockOutGet)
		r.POST("/stock-out", handler.StockOutPost)
		r.GET("/query", handler.QueryGet)
		r.POST("/query", handler.QueryPost)
		r.GET("/inventory", handler.CommodityGet)
	}
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}
