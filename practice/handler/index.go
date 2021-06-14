package handler

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/GoGin/practice/database"
	"github.com/DesistDaydream/GoGin/practice/middleware"
	"github.com/gin-gonic/gin"
)

// IndexHandler 首页界面处理
func IndexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
	fmt.Println("访问根目录后，服务端输出的信息。")
}

func IndexPOST(c *gin.Context) {
	// 用来绑定用户登录时填写的用户名和密码
	var u struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	// 将用户登陆时填写的表单数据与 u 绑定起来，以便后续使用
	c.ShouldBind(&u)

	// 处理 POST 请求的逻辑
	switch c.PostForm("button") {
	case "登录":
		// 判断用户名和密码是否正确
		userInfo, err := database.VerifyUser(u.Username, u.Password)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"err": err})
		} else {
			// 这是一个非常简化的认证方式。生成 JWT，然后将 Token 设置为 Cookie 的值。
			// 一般情况下，在前后端分离的项目中，直接将生成的 Token 返回给前端即可，具体是用 Cookie 还是用什么方式保存，由前端决定
			token, _ := middleware.GenerateToken(userInfo)
			c.SetCookie("token", token, 60, "/", "", false, true)

			// 跳转到订单页面
			c.Redirect(http.StatusFound, "/order")
		}
	case "注册":
		// 跳转到注册页面
	}
}
