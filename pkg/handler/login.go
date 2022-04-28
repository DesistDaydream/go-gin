package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Get
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
