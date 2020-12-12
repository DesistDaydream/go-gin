package handler

import (
	"net/http"

	db "github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// InventoryGet 查询页面 GET 请求处理
func InventoryGet(c *gin.Context) {
	inventory := new(db.Inventory)
	inventory.QueryData(c)
	// 页面展示处理
	h := gin.H{
		"products":    db.Products,
		"sizes":       db.Sizes,
		"amounts":     db.Amounts,
		"createTimes": db.CreateTimes,
	}
	c.HTML(http.StatusOK, "inventory.html", h)
}
