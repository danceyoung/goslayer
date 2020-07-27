package layer

type template interface {
	mainTemplate() string
	routerTemplate() string
	baseHandlerTemplate() string
	eventHandlerTemplate() string
	httpMiddlewareTemplate() string
}

func newTemplate(webframework string) template {
	return httpHandlerTemplate{}
}

type httpHandlerTemplate struct{}

func (hht httpHandlerTemplate) mainTemplate() string {
	return `package main

import (
	_ "goslayer/cmd/myapp/router"
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}`
}

func (hht httpHandlerTemplate) routerTemplate() string {
	return `package router

import (
	"goslayer/cmd/myapp/router/handler"
	"goslayer/internal/pkg/middleware"
	"net/http"
)

func init() {
	path := "/goslayer"
	eh := new(handler.EventHandler)
	http.Handle(path+"/events", middleware.HttpSet(eh.Events))
}`
}

func (hht httpHandlerTemplate) baseHandlerTemplate() string {
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

func (hht httpHandlerTemplate) eventHandlerTemplate() string {
	return `package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"goslayer/internal/myapp/event"
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

func (hht httpHandlerTemplate) httpMiddlewareTemplate() string {
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
