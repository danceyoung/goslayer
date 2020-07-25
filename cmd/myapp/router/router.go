package router

import (
	"net/http"

	"github.com/danceyoung/goslayer/cmd/myapp/router/handler"
	"github.com/danceyoung/goslayer/internal/pkg/middleware"
)

func init() {
	path := "/goslayer"
	eh := new(handler.EventHandler)
	http.Handle(path+"/events", middleware.HttpSet((eh.Events)))
}
