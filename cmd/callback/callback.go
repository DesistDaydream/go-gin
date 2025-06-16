package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"

	"github.com/DesistDaydream/go-gin/cmd/callback/baidu"
	"github.com/DesistDaydream/go-gin/config"
	logging "github.com/DesistDaydream/logging/pkg/logrus_init"
)

type Flags struct {
	ConfigFile string
}

func main() {
	var (
		logFlags logging.LogrusFlags
		flags    Flags
	)

	logging.AddFlags(&logFlags)
	pflag.StringVar(&flags.ConfigFile, "config.file", "./example/config.yaml", "配置文件路径")
	pflag.Parse()

	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败: ", err)
	}

	configFile, err := os.ReadFile(flags.ConfigFile)
	if err != nil {
		logrus.Fatal("读取配置文件失败: ", err)
	}

	if err := yaml.Unmarshal(configFile, &config.C); err != nil {
		logrus.Fatal("解析配置文件失败: ", err)
	}

	// 打印 config 中的内容
	logrus.Infof("配置文件内容: %v", config.C)

	logLevel := logrus.GetLevel()
	if logLevel != logrus.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/baidu/callback", baidu.CallBackForBaidu)

	r.Run("0.0.0.0:10000")
}

// GET
// https://openapi.baidu.com/oauth/2.0/authorize?
// response_type=code&
// client_id=您应用的AppKey&
// redirect_uri=您应用的授权回调地址&
// scope=basic,netdisk&
// device_id=您应用的AppID
