package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// HomeGet 首页界面处理
func HomeGet(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.html", nil)
	fmt.Println("访问了家目录")
}
