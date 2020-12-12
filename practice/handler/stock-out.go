package handler

import (
	"net/http"
	"regexp"

	db "github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"
)

// StockOutGet 出库页面 GET 请求处理
func StockOutGet(c *gin.Context) {
	c.HTML(http.StatusOK, "stock-out.html", nil)
}

// StockOutPost 出库页面 POST 请求处理
func StockOutPost(c *gin.Context) {
	switch c.PostForm("button") {

	// 处理出库请求
	case "出库":
		if matchResult, _ := regexp.MatchString("[1-9]+", c.PostForm("amount")); matchResult == false {
			c.String(http.StatusOK, "请填写大于0的正整数")
		} else {
			inventory := new(db.Inventory)
			inventory.AddData(c)
			c.HTML(http.StatusOK, "stock-out.html", gin.H{
				"result": "出库请求已受理！又赚钱啦！",
			})
		}

	// 返回order页面
	case "返回":
		c.Redirect(http.StatusMovedPermanently, "/order")
	}
}
