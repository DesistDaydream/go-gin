package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MySQL 数据库连接信息
type Sqlite struct {
	UserName string
	Password string
	Protocol string
	Server   string
	Port     int64
	Database string
}

// ConnDB 连接数据库
func (s *Sqlite) ConnDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
