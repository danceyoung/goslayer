package handler

import (
	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (baseh *BaseHandler) responseOk(c *gin.Context, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "ok"
	result["data"] = data
	c.JSON(200, result)
}

func (baseh *BaseHandler) recoverPanic(c *gin.Context) {
	if err := recover(); err != nil {
		baseh.responseError(c, err.(error))
	}
}

func (baseh *BaseHandler) responseError(c *gin.Context, err error) {
	result := make(map[string]interface{})
	result["code"] = -1
	result["msg"] = err.Error()
	c.JSON(200, result)
}
