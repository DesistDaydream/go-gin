package routerset

import (
	"github.com/gin-gonic/gin"
)

// RouterSet 登录页面处理
func RouterSet(route *gin.Engine) {
	route.GET("/", IndexGet)

	route.GET("/login", LoginGet)

	route.POST("/login", LoginPost)

	route.GET("/order", OrderGet)

	route.POST("/order", OrderPost)

	route.GET("/stock-in", StockInGet)

	route.POST("/stock-in", StockInPost)

	route.GET("/stock-out", StockOutGet)

	route.POST("/stock-out", StockOutPost)

	route.GET("/query", QueryGet)

	route.POST("/query", QueryPost)

	route.GET("inventory", InventoryGet)
}
