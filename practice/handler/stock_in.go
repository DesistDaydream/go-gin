package handler

import (
	"net/http"
	"regexp"

	"github.com/DesistDaydream/GoGin/practice/database"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// StockInGet 入库页面 GET 请求处理
func StockInGet(c *gin.Context) {
	c.HTML(http.StatusOK, "stock-in.gohtml", nil)
}

// StockInPost 入库页面 POST 请求处理
func StockInPost(c *gin.Context) {
	switch c.PostForm("button") {

	// 处理入库请求
	case "入库":
		if matchResult, _ := regexp.MatchString("[1-9]+", c.PostForm("amount")); matchResult == false {
			c.String(http.StatusOK, "请填写大于0的正整数")
		} else {
			i := new(database.StockInOrder)
			// 使用 gin 的绑定功能，以便接收 POST 请求中的 Form Data(表单数据)。如果不绑定，则表单数据无法写入到后端。
			c.ShouldBind(&i)

			// 调用添加订单逻辑,由于上面绑定了结构体与表单，所以表单中填写的数据将会传到结构体中。
			// 调用 AddStockInOrder() 时，也就能正常使用结构体中的数据了。
			i.AddStockInOrder(c)

			c.HTML(http.StatusOK, "stock-in.gohtml", gin.H{
				"result": "入库请求已受理！恭喜进货！",
			})
		}

	// 返回order页面
	case "返回":
		c.Redirect(http.StatusMovedPermanently, "/order")
	}
}
