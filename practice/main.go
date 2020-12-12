package main

import (
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
	// 运行 gin 程序
	r.Run()
}
