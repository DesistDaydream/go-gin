package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/go-gin/pkg/database"
	"github.com/DesistDaydream/go-gin/pkg/router"

	logging "github.com/DesistDaydream/logging/pkg/logrus_init"
)

func main() {
	var (
		logFlags logging.LogrusFlags
	)
	logging.AddFlags(&logFlags)
	pflag.Parse()

	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败: ", err)
	}

	// 初始化 gin 引擎
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("web/*/*")

	// 初始化路由
	router.InitRouter(r)

	// 设置连接数据库的信息
	// db := new(database.Sqlite)
	// db.UserName = "root"
	// db.Password = "mysql"
	// db.Protocol = "tcp"
	// db.Server = "0.0.0.0"
	// db.Port = 3306
	// db.Database = "practice"
	db := database.NewSqlite("test.db")
	// 连接数据库
	db.ConnDB()

	// 运行 gin 程序
	r.Run("0.0.0.0:8080")
}
