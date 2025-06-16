package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DesistDaydream/go-gin/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 官方文档: https://pan.baidu.com/union/doc/ol0rsap9s

const oauthURL = "https://openapi.baidu.com/oauth/2.0/token"

type BaiduAccessToken struct {
	ExpiresIn     int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
	SessionKey    string `json:"session_key"`
	Scope         string `json:"scope"`
}

func CallBackForBaidu(c *gin.Context) {
	code := c.Query("code")
	logrus.Debugf("授权码: %v", code)

	client := &http.Client{}
	req, err := http.NewRequest("GET", oauthURL, nil)
	if err != nil {
		logrus.Warn(err)
		return
	}

	q := req.URL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("code", code)
	q.Add("client_id", config.C.Callback.Baidu.ClientID)
	q.Add("client_secret", config.C.Callback.Baidu.ClientSecret)
	q.Add("redirect_uri", config.C.Callback.Baidu.RedirectURI)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	var accessToken BaiduAccessToken
	if err := json.Unmarshal(resBody, &accessToken); err != nil {
		logrus.Error(err)
		return
	}

	logrus.Debugf("accessToken: %v", accessToken)

	c.JSON(http.StatusOK, accessToken)
}

func RefreshToken(refreshToken string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", oauthURL, nil)
	if err != nil {
		logrus.Warn(err)
		return
	}

	q := req.URL.Query()
	q.Add("grant_type", "refresh_token")
	q.Add("refresh_token", refreshToken)
	q.Add("client_id", config.C.Callback.Baidu.ClientID)
	q.Add("client_secret", config.C.Callback.Baidu.ClientSecret)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	var accessToken BaiduAccessToken
	if err := json.Unmarshal(resBody, &accessToken); err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("accessToken: %v\n", string(resBody))
}
