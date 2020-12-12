package handler

import (
	"fmt"
	"net/http"

	db "github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// QueryGet 查询页面 GET 请求处理
func QueryGet(c *gin.Context) {
	inventory := new(db.Inventory)
	inventory.QueryData(c)
	// 页面展示处理
	h := gin.H{
		"products":    db.Products,
		"sizes":       db.Sizes,
		"amounts":     db.Amounts,
		"createTimes": db.CreateTimes,
	}
	c.HTML(http.StatusOK, "query.html", h)
}

// QueryPost 查询页面 POST 请求处理
func QueryPost(c *gin.Context) {
	switch c.PostForm("button") {
	case "查询":
		c.Redirect(http.StatusMovedPermanently, "/inventory")
	case "返回":
		c.Redirect(http.StatusMovedPermanently, "/order")
	}
	fmt.Println("显示当前库存数：")
}
