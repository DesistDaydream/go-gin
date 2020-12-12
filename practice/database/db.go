package database

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// Inventory 产品库存的信息，后面的描述信息用来进行数据绑定
type Inventory struct {
	Product    string `form:"product" binding:"required"`
	Size       string `form:"size" binding:"required"`
	Amount     int    `form:"amount" binding:"required"`
	CreateTime string
}

var (
	// InverntoryBind 用于将form中传递的参数与Invertory结构体中的属性绑定
	InverntoryBind Inventory
	// Products 产品集合，用于在前端遍历数据并逐行展示
	Products []string
	// Sizes 尺寸集合，用于在前端遍历数据并逐行展示
	Sizes []string
	// Amounts 库存集合，用于在前端遍历数据并逐行展示
	Amounts []int
	// CreateTimes 入库时间集合，用于在前端遍历数据并逐行展示
	CreateTimes []string
	// DB 123
	db *sql.DB
)

// 数据库连接
func dbConn() {
	var err error
	db, err = sql.Open("mysql", "root:mypassword@tcp(10.10.100.200:3306)/caredaily?charset=utf8")
	CheckErr(err)
}

// AddData 在stock-in.go中添加向数据库添加数据
func (i *Inventory) AddData(c *gin.Context) {
	fmt.Println("入库类型：", c.PostForm("product"))
	fmt.Println("入库尺寸：", c.PostForm("size"))
	fmt.Println("入库数量：", c.PostForm("amount"))

	// 使用 gin 的绑定功能，将 Inventory 结构体中的属性与表单传入的参数绑定
	c.ShouldBind(&InverntoryBind)

	// 数据处理
	dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT inventory SET product=?,size=?,amount=?")
	CheckErr(err)

	switch c.PostForm("button") {
	case "入库":
		_, err := stmt.Exec(InverntoryBind.Product, InverntoryBind.Size, InverntoryBind.Amount)
		CheckErr(err)
	case "出库":
		_, err := stmt.Exec(InverntoryBind.Product, InverntoryBind.Size, -InverntoryBind.Amount)
		CheckErr(err)
	}
}

// QueryData 在query.go中查询数据库
func (i *Inventory) QueryData(c *gin.Context) {
	Products = make([]string, 0)
	Sizes = make([]string, 0)
	Amounts = make([]int, 0)
	CreateTimes = make([]string, 0)

	// 数据处理
	dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT product,size,amount,create_time FROM inventory")
	CheckErr(err)

	for rows.Next() {

		err = rows.Scan(&i.Product, &i.Size, &i.Amount, &i.CreateTime)
		CheckErr(err)

		Products = append(Products, i.Product)
		Sizes = append(Sizes, i.Size)
		Amounts = append(Amounts, i.Amount)
		CreateTimes = append(CreateTimes, i.CreateTime)

		fmt.Println("数据库中的数据为：", *i)
	}
}

// DelData 在stock-in.go中删除数据
func (i *Inventory) DelData(c *gin.Context) {
	dbConn()
	defer db.Close()
}

// CheckErr 检查数据库操作
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
