package main

import (
	"github.com/danceyoung/goslayer/cmd/gmyapp/router/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/goslayer/events", (&handler.EventHandler{}).Events)
	r.POST("goslayer/events/join", (&handler.EventHandler{}).JoinAEvent)
	r.Run(":8080")
}
