package handler

import (
	"io"
	"net/http"
)

type EventHandler struct {
	BaseHandler
}

func (eventh *EventHandler) Events(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello World")
}
