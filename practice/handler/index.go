package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGet 首页界面处理
func IndexGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
	fmt.Println("访问根目录后，服务端输出的信息。")
}
