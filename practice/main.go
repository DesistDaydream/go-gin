package main

import (
	"github.com/DesistDaydream/GoGin/practice/database"
	"github.com/DesistDaydream/GoGin/practice/router"

	"github.com/gin-gonic/gin"
)

// var route *gin.Engine

func main() {
	// 初始化 gin 引擎
	r := gin.Default()

	// 加载模板文件
	r.LoadHTMLGlob("templates/*")

	// 初始化路由
	router.InitRouter(r)

	// 设置连接数据库的信息
	c := new(database.ConnDatabaseInfo)
	c.UserName = "root"
	c.Password = "mysql"
	c.Protocol = "tcp"
	c.Server = "0.0.0.0"
	c.Port = 3306
	c.Database = "practice"
	// 连接数据库
	c.ConnDB()

	// 运行 gin 程序
	r.Run()
}
