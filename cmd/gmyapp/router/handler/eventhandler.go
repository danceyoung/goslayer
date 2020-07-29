package handler

import (
	"fmt"

	"github.com/danceyoung/goslayer/internal/myapp/event"
	"github.com/gin-gonic/gin"
)

type EventHandler struct{}

func (eventh *EventHandler) Events(c *gin.Context) {
	events, _ := event.Events()
	c.JSON(200, gin.H{"data": events})
}

func (eventh *EventHandler) JoinAEvent(c *gin.Context) {
	var m event.Member
	c.ShouldBindJSON(&m)
	fmt.Println("dd ", m)
	event.JoinAEvent(c.Query("event-id"), m)
}
