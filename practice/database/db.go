package database

import (
	"fmt"
	"log"

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

// Commodity 一个商品应该具有的属性，后面的描述信息用来进行数据绑定
type Commodity struct {
	Product    string `form:"product" binding:"required"`
	Size       string `form:"size" binding:"required"`
	Amount     int    `form:"amount" binding:"required"`
	CreateTime string
}

var (
	// Products 产品集合，用于在前端遍历数据并逐行展示
	Products []string
	// Sizes 尺寸集合，用于在前端遍历数据并逐行展示
	Sizes []string
	// Amounts 库存集合，用于在前端遍历数据并逐行展示
	Amounts []int
	// CreateTimes 入库时间集合，用于在前端遍历数据并逐行展示
	CreateTimes []string
	db          *gorm.DB
)

// ConnDB 连接数据库
func ConnDB() {
	var err error
	db, err = gorm.Open("mysql", "root:mysql@tcp(0.0.0.0:3306)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("failed to connect database, ", err)
	}
	// 刷新数据表，
	db.AutoMigrate(&Commodity{})
}

// AddData 在stock-in.go中添加向数据库添加数据
func (com *Commodity) AddData(c *gin.Context) {
	fmt.Println("入库类型：", c.PostForm("product"))
	fmt.Println("入库尺寸：", c.PostForm("size"))
	fmt.Println("入库数量：", c.PostForm("amount"))

	// 使用 gin 的绑定功能，将 Commodity 结构体中的属性与表单传入的参数绑定
	var data Commodity
	c.ShouldBind(&data)

	// 数据处理
	ConnDB()
	defer db.Close()

	switch c.PostForm("button") {
	case "入库":
		db.Create(data)
	case "出库":
	}
}

// QueryData 在 query.go 中查询数据库
func (com *Commodity) QueryData(c *gin.Context) {
	// 重置数据，避免后面的查询包含前一次查询内容
	Products = make([]string, 0)
	Sizes = make([]string, 0)
	Amounts = make([]int, 0)
	CreateTimes = make([]string, 0)
	// 数据处理
	ConnDB()
	defer db.Close()

	var Commodities []Commodity
	db.Find(&Commodities)
	for _, commodity := range Commodities {
		Products = append(Products, commodity.Product)
		Sizes = append(Sizes, commodity.Size)
		Amounts = append(Amounts, commodity.Amount)
	}

	fmt.Println("数据库中的数据为：", Commodities)
}

// DelData 在stock-in.go中删除数据
func (com *Commodity) DelData(c *gin.Context) {
	ConnDB()
	defer db.Close()
}

// CheckErr 检查数据库操作
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
