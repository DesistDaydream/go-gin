package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL 数据库连接信息
type MySQL struct {
	UserName string
	Password string
	Protocol string
	Server   string
	Port     int64
	Database string
}

func NewMySQL(userName, password, protocol, server string, port int64, database string) *MySQL {
	return &MySQL{
		UserName: userName,
		Password: password,
		Protocol: protocol,
		Server:   server,
		Port:     port,
		Database: database,
	}
}

// ConnDB 连接数据库
func (c *MySQL) ConnDB() {
	log.Println("开始连接数据库")
	connInfo := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.UserName, c.Password, c.Protocol, c.Server, c.Port, c.Database)
	log.Println("数据库连接信息：", connInfo)
	db, err = gorm.Open(mysql.Open(connInfo), &gorm.Config{})
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
