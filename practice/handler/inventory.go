package handler

import (
	"net/http"

	"github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// CommodityGet 查询页面 GET 请求处理
func CommodityGet(c *gin.Context) {
	order := new(database.StockInOrder)
	order.QueryStockInOrder(c)
	// 页面展示处理
	h := gin.H{
		"products": database.Products,
		"sizes":    database.Sizes,
		"amounts":  database.Amounts,
	}
	c.HTML(http.StatusOK, "inventory.html", h)
}
