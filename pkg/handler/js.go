package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.js", nil)
}
