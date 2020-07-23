package router

import (
	"net/http"

	"github.com/danceyoung/goslayer/cmd/myapp/router/handler"
)

func init() {
	path := "/goslayer"
	eh := new(handler.EventHandler)
	http.Handle(path+"/events", http.HandlerFunc(eh.Events))
}
