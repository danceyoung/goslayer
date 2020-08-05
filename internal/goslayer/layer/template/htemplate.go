package template

// Template is a interface for providing some template func
type Template interface {
	MainTemplate() string
	RouterTemplate() string
	BaseHandlerTemplate() string
	EventHandlerTemplate() string
	HttpMiddlewareTemplate() string
	EventBizTemplate() string
	PkgDbMysqlTemplate() string
}

// HttpHandlerTemplate implements Template and provides some go files base on http.Handler buildin
type HttpHandlerTemplate struct {
	baseTemplate
}

// MainTemplate provides content for main.go file
func (hhtmpl HttpHandlerTemplate) MainTemplate() string {
	return `package main

import (
	"log"
	"net/http"

	_ "github.com/danceyoung/goslayer/cmd/myapp/router"
)

func main() {
	log.Println("[http Handler] Listening and serving HTTP on :8080")
	log.Println(http.ListenAndServe(":8080", nil))
}`
}

// RouterTemplate provides content for router.go file
func (hhtmpl HttpHandlerTemplate) RouterTemplate() string {
	return `package router

import (
	"net/http"

	"github.com/danceyoung/goslayer/cmd/myapp/router/handler"
	"github.com/danceyoung/goslayer/internal/pkg/middleware"
)

func init() {
	path := "/goslayer"

	http.Handle(path+"/events", middleware.HttpSet(((&handler.EventHandler{}).Events)))
	http.Handle(path+"/events/join", middleware.HttpSet(((&handler.EventHandler{}).JoinAEvent)))
}`
}

// BaseHandlerTemplate provides content for basehandler.go file
func (hhtmpl HttpHandlerTemplate) BaseHandlerTemplate() string {
	return `package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type BaseHandler struct{}

func (baseh *BaseHandler) responseOk(rw http.ResponseWriter, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "ok"
	result["data"] = data
	bytes, _ := json.Marshal(result)
	io.WriteString(rw, string(bytes))
}

func (baseh *BaseHandler) recoverPanic(rw http.ResponseWriter) {
	if err := recover(); err != nil {
		baseh.responseError(rw, err.(error))
	}
}

func (baseh *BaseHandler) responseError(rw http.ResponseWriter, err error) {
	result := make(map[string]interface{})
	result["code"] = -1
	result["msg"] = err.Error()
	bytes, _ := json.Marshal(result)
	io.WriteString(rw, string(bytes))
}`
}

// EventHandlerTemplate provides content for eventhandler.go file
func (hhtmpl HttpHandlerTemplate) EventHandlerTemplate() string {
	return `package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/danceyoung/goslayer/internal/myapp/event"
)

type EventHandler struct {
	BaseHandler
}

func (eventh *EventHandler) Events(rw http.ResponseWriter, req *http.Request) {
	events, err := event.Events()
	if err != nil {
		eventh.responseError(rw, err)
	}

	eventh.responseOk(rw, events)
}

func (eventh *EventHandler) JoinAEvent(rw http.ResponseWriter, req *http.Request) {
	defer eventh.recoverPanic(rw)
	req.ParseForm()

	bodybytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	if req.Method != http.MethodPost {
		panic(errors.New("not matched handler"))
	}

	var m event.Member
	err = json.Unmarshal(bodybytes, &m)
	if err != nil {
		panic(err)
	}

	err = event.JoinAEvent(req.Form.Get("event-id"), m)
	if err != nil {
		panic(err)
	}

	eventh.responseOk(rw, nil)
}`
}

// HttpMiddlewareTemplate provides content for httpset.go file
func (hhtmpl HttpHandlerTemplate) HttpMiddlewareTemplate() string {
	return `package middleware

import (
	"log"
	"net/http"
)

func HttpSet(hf func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.Method+" the url requesting is ", req.URL)
		rw.Header().Set("Content-Type", "application/json")
		hf(rw, req)
	})
}`
}

// EventBizTemplate provides content for business logic
func (hhtmpl HttpHandlerTemplate) EventBizTemplate() string {
	return hhtmpl.eventbizTemplate()
}

// PkgDbMysqlTemplate provides content for mysql.go file insider internal/pkg/db
func (hhtmpl HttpHandlerTemplate) PkgDbMysqlTemplate() string {
	return hhtmpl.pkgdbmysqlTemplate()
}
