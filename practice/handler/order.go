package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderGet 订单页面 GET 请求处理
func OrderGet(c *gin.Context) {
	c.HTML(http.StatusOK, "order.html", gin.H{"title": "订单管理系统"})
}

// OrderPost 订单页面 POST 请求处理
func OrderPost(c *gin.Context) {
	switch c.PostForm("button") {
	case "入库":
		c.Redirect(http.StatusFound, "/stock-in")
	case "出库":
		c.Redirect(http.StatusFound, "/stock-out")
	case "查询":
		c.Redirect(http.StatusFound, "/query")
	}
}
