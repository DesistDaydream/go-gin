package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnDatabaseInfo 数据库连接信息
type ConnDatabaseInfo struct {
	UserName string
	Password string
	Protocol string
	Server   string
	Port     int64
	Database string
}

var (
	err error
	db  *gorm.DB
)

// ConnDB 连接数据库
func (c *ConnDatabaseInfo) ConnDB() {
	// fmt.Println("开始连接数据库")
	// connInfo := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.UserName, c.Password, c.Protocol, c.Server, c.Port, c.Database)
	// fmt.Println("数据库连接信息：", connInfo)
	// db, err = gorm.Open("mysql", connInfo)
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database, ", err)
	}

	// db.DB().SetConnMaxLifetime(60 * time.Second)

	// 刷新数据表
	db.AutoMigrate(&StockInOrder{}, &User{})

	// 创建管理员用户
	if err := createAdminUser(); err != nil {
		log.Fatalln(err)
	}
}
