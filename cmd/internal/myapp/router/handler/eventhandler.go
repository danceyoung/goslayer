package handler

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
}
