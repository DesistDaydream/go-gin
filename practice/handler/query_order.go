package handler

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// QueryGet 查询页面 GET 请求处理
func QueryGet(c *gin.Context) {
	// 重置数据，避免后面的查询包含前一次查询内容
	// 各种数据的数组，用于在前端遍历数据并逐行展示
	Providers := make([]string, 0)
	Products := make([]string, 0)
	Sizes := make([]string, 0)
	Amounts := make([]int, 0)

	for _, order := range database.QueryStockInOrder() {
		Providers = append(Providers, order.Provider)
		Products = append(Products, order.Product)
		Sizes = append(Sizes, order.Size)
		Amounts = append(Amounts, order.Amount)
	}

	// 页面展示处理
	h := gin.H{
		"provider": Providers,
		"products": Products,
		"sizes":    Sizes,
		"amounts":  Amounts,
	}
	c.HTML(http.StatusOK, "query.gohtml", h)
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
