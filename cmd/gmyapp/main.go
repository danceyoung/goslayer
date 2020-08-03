package main

import (
	"github.com/danceyoung/goslayer/cmd/gmyapp/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)
	r.Run(":8080")
}
