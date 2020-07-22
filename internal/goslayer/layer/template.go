package layer

type template interface {
	mainTemplate() string
	routerTemplate() string
	baseHandlerTemplate() string
	eventHandlerTemplate() string
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
	"net/http"
)

func init() {
	path := "/goslayer"
	eh := new(handler.EventHandler)
	http.Handle(path+"/events", http.HandlerFunc(eh.Events))
}`
}

func (hht httpHandlerTemplate) baseHandlerTemplate() string {
	return `package handler

type BaseHandler struct{}`
}

func (hht httpHandlerTemplate) eventHandlerTemplate() string {
	return `package handler

import (
	"io"
	"net/http"
)

type EventHandler struct {
	BaseHandler
}

func (eventh *EventHandler) Events(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello World")
}`
}
