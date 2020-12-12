package handler

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

func handleData() (Providers []string, Products []string, Sizes []string, Amounts []int) {
	// 调用查询逻辑，获取数据，并处理数据以便展示
	for _, order := range database.QueryStockInOrder() {
		Providers = append(Providers, order.Provider)
		Products = append(Products, order.Product)
		Sizes = append(Sizes, order.Size)
		Amounts = append(Amounts, order.Amount)
	}
	return
}

// QueryGet 查询页面 GET 请求处理
func QueryGet(c *gin.Context) {
	// 从数据库中获取数据，订单的每个属性都是一个数组，以便可以轮询写入到前端展示页面的表格中
	Providers, Products, Sizes, Amounts := handleData()

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
