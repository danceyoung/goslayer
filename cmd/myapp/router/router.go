package router

import (
	"goslayer/cmd/myapp/router/handler"
	"net/http"
)

func init() {
	path := "/goslayer"
	eh := new(handler.EventHandler)
	http.Handle(path+"/events", http.HandlerFunc(eh.Events))
}
