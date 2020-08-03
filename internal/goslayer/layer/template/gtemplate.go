package template

type GINTemplate struct{}

func (gt GINTemplate) MainTemplate() string {
	return `package main

import (
	"github.com/danceyoung/goslayer/cmd/myapp/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)
	r.Run(":8080")
}`
}
func (gt GINTemplate) RouterTemplate() string {
	return `package router

import (
	"github.com/danceyoung/goslayer/cmd/myapp/router/handler"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/goslayer/events", (&handler.EventHandler{}).Events)
	engine.POST("goslayer/events/join", (&handler.EventHandler{}).JoinAEvent)
}`
}
func (gt GINTemplate) BaseHandlerTemplate() string {
	return `package handler

import (
	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (baseh *BaseHandler) responseOk(c *gin.Context, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "ok"
	result["data"] = data
	c.JSON(200, result)
}

func (baseh *BaseHandler) recoverPanic(c *gin.Context) {
	if err := recover(); err != nil {
		baseh.responseError(c, err.(error))
	}
}

func (baseh *BaseHandler) responseError(c *gin.Context, err error) {
	result := make(map[string]interface{})
	result["code"] = -1
	result["msg"] = err.Error()
	c.JSON(200, result)
}`
}
func (gt GINTemplate) EventHandlerTemplate() string {
	return `package handler

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
}`
}
func (gt GINTemplate) HttpMiddlewareTemplate() string {
	return ``
}
func (gt GINTemplate) EventBizTemplate() string {
	return `package event

import "errors"

//implement biz logic and wrap response data
func Events() ([]map[string]interface{}, error) {
	return events(), nil
}

//query events from db,eg:mysql
func events() []map[string]interface{} {
	var result []map[string]interface{}
	result = append(result, map[string]interface{}{"id": 1, "event_name": "dancing competition"}, map[string]interface{}{"id": 1, "event_name": "singing competition"})
	return result

}

type Member struct {
	Name  string
	Email string
}

func JoinAEvent(eventid string, member Member) error {
	if len(eventid) == 0 || len(member.Name) == 0 || len(member.Email) == 0 {
		return errors.New("parmas are not enough")
	}
	if err := joinAEvent(eventid, member); err != nil {
		return errors.New("join a event occurring a error: " + err.Error())
	}
	return nil
}

//insert a record into db
func joinAEvent(eventid string, member Member) error {
	return nil
}`
}
