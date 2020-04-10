package handler

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"ddd/infrastructure/common/responseCode/code"
)

// Response 返回结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseCoder 返回码
type ResponseCoder interface {
	String() string
}

// SendResponse 返回结构
func SendResponse(c *gin.Context, coder ResponseCoder, data interface{}) {
	code, message := decode(coder)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// DecodeErr 解析返回码
func decode(coder ResponseCoder) (int, string) {
	if coder == nil {
		return int(code.Code_OK), code.Code_OK.String()
	}

	intCode := reflect.ValueOf(coder).Int()
	return int(intCode), coder.String()
}
