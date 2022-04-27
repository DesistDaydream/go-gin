package handler

import (
	"net/http"

	"github.com/DesistDaydream/go-gin/pkg/database"
	"github.com/DesistDaydream/go-gin/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Code    int    `json:"code"`
	Token   string `json:"token"`
	Message string `json:"msg"`
}

func Login(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 用来绑定用户登录时填写的用户名和密码
	var u struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	// 将用户登陆时填写的表单数据与 u 绑定起来，以便后续使用
	c.ShouldBind(&u)

	// 判断用户名和密码是否正确
	if userInfo, err := database.VerifyUser(u.Username, u.Password); err != nil {
		// 设置响应体
		resp := LoginResponse{
			Code:    0,
			Message: "用户名或密码错误",
			Token:   "",
		}
		c.JSON(http.StatusOK, resp)
	} else {
		// 若用户名/密码验证成功，则生成 JWT，并以 Cookie 的形式交给客户端。
		token, _ := middleware.GenerateToken(userInfo)
		c.SetCookie("token", token, 60, "/", "", false, true)

		// 设置响应体
		resp := LoginResponse{
			Code:    1,
			Message: "登录成功",
			Token:   token,
		}
		// 返回 Token 到前端
		c.JSON(http.StatusOK, resp)
	}
}
