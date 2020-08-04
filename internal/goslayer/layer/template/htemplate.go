package template

type Template interface {
	MainTemplate() string
	RouterTemplate() string
	BaseHandlerTemplate() string
	EventHandlerTemplate() string
	HttpMiddlewareTemplate() string
	EventBizTemplate() string
}

type HttpHandlerTemplate struct{}

func (hht HttpHandlerTemplate) MainTemplate() string {
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

func (hht HttpHandlerTemplate) RouterTemplate() string {
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

func (hht HttpHandlerTemplate) BaseHandlerTemplate() string {
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

func (hht HttpHandlerTemplate) EventHandlerTemplate() string {
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

func (hht HttpHandlerTemplate) HttpMiddlewareTemplate() string {
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

func (hht HttpHandlerTemplate) EventBizTemplate() string {
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
