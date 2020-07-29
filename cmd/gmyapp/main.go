package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/goslayer/events", func(context *gin.Context) { fmt.Println("hello gin") })
	r.Run(":8080")
}
