package database

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DatabaseInfo 数据库连接信息
type DatabaseInfo struct {
	UserName string
	Password string
	Protocol string
	Server   string
	Port     int64
	Database string
}

// Order 一个订单的属性，后面的描述信息用来绑定属性与表单中的字段
type Order struct {
	// gorm.Model
	Provider string `form:"provider" binding:"required"`
	Commodity
}

// Commodity 一个商品应该具有的属性
type Commodity struct {
	Product string `form:"product" binding:"required"`
	Size    string `form:"size" binding:"required"`
	Amount  int    `form:"amount" binding:"required"`
}

var (
	// Providers 供应商集合，用于在前端遍历数据并逐行展示
	Providers []string
	// Products 产品集合，用于在前端遍历数据并逐行展示
	Products []string
	// Sizes 尺寸集合，用于在前端遍历数据并逐行展示
	Sizes []string
	// Amounts 库存集合，用于在前端遍历数据并逐行展示
	Amounts []int
	db      *gorm.DB
)

// AddData 在 stock-in.go 中向数据库添加数据
func (o *Order) AddData(c *gin.Context) {
	fmt.Println("供应商", c.PostForm("provider"))
	fmt.Println("入库类型：", c.PostForm("product"))
	fmt.Println("入库尺寸：", c.PostForm("size"))
	fmt.Println("入库数量：", c.PostForm("amount"))

	// 使用 gin 的绑定功能，将 Commodity 结构体中的属性与表单传入的参数绑定，以便将表单的值应用到结构体中，不绑定的话，入库数据为空
	c.ShouldBind(&o)

	// 数据处理
	ConnDB()
	defer db.Close()

	switch c.PostForm("button") {
	case "入库":
		db.Create(o)
	case "出库":
	}
}

// QueryData 在 query.go 中查询数据库
func (o *Order) QueryData(c *gin.Context) {
	// 重置数据，避免后面的查询包含前一次查询内容
	Providers = make([]string, 0)
	Products = make([]string, 0)
	Sizes = make([]string, 0)
	Amounts = make([]int, 0)
	// 数据处理
	ConnDB()
	defer db.Close()

	var Orders []Order
	db.Find(&Orders)
	for _, order := range Orders {
		Providers = append(Providers, order.Provider)
		Products = append(Products, order.Product)
		Sizes = append(Sizes, order.Size)
		Amounts = append(Amounts, order.Amount)
	}

	fmt.Println("数据库中的数据为：", Orders)
}

// DelData 在stock-in.go中删除数据
func (o *Order) DelData(c *gin.Context) {
	ConnDB()
	defer db.Close()
}
