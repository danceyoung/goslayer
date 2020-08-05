package template

// GINTemplate implements Template and provides some go files base on http.Handler building
type GINTemplate struct {
	baseTemplate
}

// MainTemplate provides content for main.go file
func (gintmpl GINTemplate) MainTemplate() string {
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

// RouterTemplate provides content for router.go file
func (gintmpl GINTemplate) RouterTemplate() string {
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

// BaseHandlerTemplate provides content for basehandler.go file
func (gintmpl GINTemplate) BaseHandlerTemplate() string {
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

// EventHandlerTemplate provides content for eventhandler.go file
func (gintmpl GINTemplate) EventHandlerTemplate() string {
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

// HttpMiddlewareTemplate provides content for middleware, but here is not implement
func (gintmpl GINTemplate) HttpMiddlewareTemplate() string {
	return ``
}

// EventBizTemplate provides content for business logic
func (gintmpl GINTemplate) EventBizTemplate() string {
	return gintmpl.eventbizTemplate()
}

// PkgDbMysqlTemplate provides content for mysql.go file insider internal/pkg/db
func (gintmpl GINTemplate) PkgDbMysqlTemplate() string {
	return gintmpl.pkgdbmysqlTemplate()
}
