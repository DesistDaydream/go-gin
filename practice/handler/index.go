package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler 首页界面处理
func IndexHandler(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		c.Redirect(http.StatusFound, "login.html")
	default:
		c.HTML(http.StatusOK, "index.html", nil)
		fmt.Println("访问根目录后，服务端输出的信息。")
	}
}
