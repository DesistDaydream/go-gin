package handler

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/GoGin/practice/database"
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
		_, err := database.VerifyUser(u.Username, u.Password)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"err": err})
		} else {
			// TODO：
			// 登录成功，在当前这个用户得 SessionData 保存一个键值对 isLogin=true
			// 先从上下文中获取 SessionData
			// 给 SessionData 设置 isLogin=true
			//
			// 跳转到订单页面
			c.Redirect(http.StatusFound, "/order")
		}
	case "注册":
		// 跳转到注册页面
	}
}
