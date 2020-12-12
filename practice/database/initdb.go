package database

import (
	"fmt"
	"log"

	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// ConnDB 连接数据库
func ConnDB() {
	fmt.Println("开始连接数据库")
	var err error
	db, err = gorm.Open("mysql", "root:mysql@tcp(0.0.0.0:3306)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("failed to connect database, ", err)
	}
	// 刷新数据表
	db.AutoMigrate(&Order{})
}
