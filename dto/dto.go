package dto

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 约定code
// 0: 失败
// 1: token过期，前端用存储的账号密码重新登录/跳回登录页
// 2: token错误，前端清除存储的账号密码，跳回登录页
// 100: 成功

type Response struct {
	Message string `json:"message"`
	Code	int `json:"code"`
	Data    interface{} `json:"data"`
}

type ResponseOptions func(*Response)

func WithMessage(message string) ResponseOptions {
	return func(r *Response) {
		r.Message = message
	}
}

func WithCode(code int) ResponseOptions {
	return func(r *Response) {
		r.Code = code
	}
}

func WithData(data interface{}) ResponseOptions {
	return func(r *Response) {
		r.Data = data
	}
}

func SuccessResponse(c *gin.Context, opts ...ResponseOptions) {
	response := Response{
		Message: "success",
		Code: 100,
		Data: nil,
	}
	for _, opt := range opts {
		opt(&response)
	}
	c.JSON(200, response)
	log.Println(response.Data)
}

func ErrorResponse(c *gin.Context, opts ...ResponseOptions) {
	response := Response{
		Message: "error",
		Code: 0,
		Data: nil,
	}
	for _, opt := range opts {
		opt(&response)
	}
	c.JSON(500, response)
}