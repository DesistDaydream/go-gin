package main

import (
	"os"

	"github.com/DesistDaydream/GoGin/practice/database"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/GoGin/practice/router"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LogInit 日志功能初始化，若指定了 log-output 命令行标志，则将日志写入到文件中
func LogInit(level, file string) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(l)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

func main() {
	// 日志相关命令行标志
	logLevel := pflag.String("log-level", "debug", "The logging level:[debug, info, warn, error, fatal]")
	logFile := pflag.String("log-output", "", "the file which log to, default stdout")
	pflag.Parse()
	// 初始化日志
	if err := LogInit(*logLevel, *logFile); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}

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
