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

func (hht httpHandlerTemplate) httpMiddlewareTemplate() string {
	return `package middleware

import (
	"log"
	"net/http"
)

func HttpSet(hf func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println("The url requesting is ", req.URL)
		rw.Header().Set("Content-Type", "application/json")
		hf(rw, req)
	})
}`
}
