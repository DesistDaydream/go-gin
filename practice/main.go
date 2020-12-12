package main

import (
	"github.com/DesistDaydream/GoGin/practice/routerset"

	"github.com/gin-gonic/gin"
)

// var route *gin.Engine

func main() {
	route := gin.Default()

	route.LoadHTMLGlob("templates/*")

	routerset.RouterSet(route)

	route.Run()
}
