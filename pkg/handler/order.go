package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderHandler 订单页面 POST 请求处理
func OrderHandler(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		switch c.PostForm("button") {
		case "入库":
			c.Redirect(http.StatusFound, "/stock-in")
		case "出库":
			c.Redirect(http.StatusFound, "/stock-out")
		case "查询":
			c.Redirect(http.StatusFound, "/query")
		}
	default:
		c.HTML(http.StatusOK, "order.html", gin.H{"title": "订单管理系统"})
	}
}
