package baidu

import (
	"os"
	"testing"

	"github.com/DesistDaydream/go-gin/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func TestRefreshToken(t *testing.T) {
	configFile, err := os.ReadFile("../../../example/config.yaml")
	if err != nil {
		logrus.Fatal("读取配置文件失败: ", err)
	}

	if err := yaml.Unmarshal(configFile, &config.C); err != nil {
		logrus.Fatal("解析配置文件失败: ", err)
	}

	refreshToken := ""
	RefreshToken(refreshToken)
}
