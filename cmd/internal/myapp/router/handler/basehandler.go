package handler

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
}
