package router

import (
	"net/http"

	"github.com/danceyoung/goslayer/cmd/hmyapp/router/handler"
	"github.com/danceyoung/goslayer/internal/pkg/middleware"
)

func init() {
	path := "/goslayer"

	http.Handle(path+"/events", middleware.HttpSet(((&handler.EventHandler{}).Events)))
	http.Handle(path+"/events/join", middleware.HttpSet(((&handler.EventHandler{}).JoinAEvent)))
}
