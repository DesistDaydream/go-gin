package handler

import (
	"net/http"
	"regexp"

	db "github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// StockInGet 入库页面 GET 请求处理
func StockInGet(c *gin.Context) {
	c.HTML(http.StatusOK, "stock-in.html", nil)
}

// StockInPost 入库页面 POST 请求处理
func StockInPost(c *gin.Context) {
	switch c.PostForm("button") {

	// 处理入库请求
	case "入库":
		if matchResult, _ := regexp.MatchString("[1-9]+", c.PostForm("amount")); matchResult == false {
			c.String(http.StatusOK, "请填写大于0的正整数")
		} else {
			inventory := new(db.Inventory)
			inventory.AddData(c)
			c.HTML(http.StatusOK, "stock-in.html", gin.H{
				"result": "入库请求已受理！恭喜进货！",
			})
		}

	// 返回order页面
	case "返回":
		c.Redirect(http.StatusMovedPermanently, "/order")
	}
}
