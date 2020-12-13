package main

import (
	"github.com/DesistDaydream/GoGin/practice/database"

	"github.com/DesistDaydream/GoGin/practice/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 gin 引擎
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")

	// 初始化路由
	router.InitRouter(r)

	// 设置连接数据库的信息
	db := new(database.ConnDatabaseInfo)
	db.UserName = "root"
	db.Password = "mysql"
	db.Protocol = "tcp"
	db.Server = "0.0.0.0"
	db.Port = 3306
	db.Database = "practice"
	// 连接数据库
	db.ConnDB()

	// 运行 gin 程序
	r.Run()
}
