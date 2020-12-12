package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// StockInOrder 入库订单的属性，后面的描述信息用来绑定属性与表单中的字段
type StockInOrder struct {
	// gorm.Model
	Provider string `form:"provider" binding:"required"`
	Commodity
}

// StockOutOrder 出库订单的属性，后面的描述信息用来绑定属性与表单中的字段
type StockOutOrder struct {
	// gorm.Model
	Customer string `form:"customer" binding:"required"`
	Commodity
}

// Commodity 一个商品应该具有的属性
type Commodity struct {
	Product string `form:"product" binding:"required"`
	Size    string `form:"size" binding:"required"`
	Amount  int    `form:"amount" binding:"required"`
}

// AddStockInOrder 在 stock-in.go 页面中向数据库添加入库订单数据
func (i *StockInOrder) AddStockInOrder(c *gin.Context) {
	// 从后台检查客户端 POST 请求中的数据
	fmt.Printf("供应商：%v\n入库类型：%v\n入库尺寸：%v\n入库数量：%v\n", c.PostForm("provider"), c.PostForm("product"), c.PostForm("size"), c.PostForm("amount"))

	// 数据库处理逻辑，结构体中的数据在调用本方法之前已经处理好了。
	db.Create(i)

}

// QueryStockInOrder 在 query.go 页面中查询数据库中的入库订单数据
func QueryStockInOrder() (StockInOrders []StockInOrder) {
	// 数据库处理逻辑
	db.Find(&StockInOrders)

	// 在后台打印查询的数据，检查数据正确性
	fmt.Println("数据库中的数据为：", StockInOrders)
	return
}

// AddStockOutOrder 在 stock-in.go 页面添加出库订单数据
func (i *StockOutOrder) AddStockOutOrder(c *gin.Context) {
	fmt.Println("待更新")
}
