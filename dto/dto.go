package dto

import (
	"github.com/gin-gonic/gin"

	"my_diary/constant"
)

// 约定code
// 0: 失败
// 1: token过期，前端用存储的账号密码重新登录/跳回登录页
// 2: token错误，前端清除存储的账号密码，跳回登录页
// 100: 成功

type Response struct {
	Message string `json:"message"`
	Code	constant.ResponseCode `json:"code"`
	Data    interface{} `json:"data"`
}

type ResponseOptions func(*Response)

func WithMessage(message string) ResponseOptions {
	return func(r *Response) {
		r.Message = message
	}
}

func WithCode(code constant.ResponseCode) ResponseOptions {
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
		Code: constant.Success,
		Data: nil,
	}
	for _, opt := range opts {
		opt(&response)
	}
	c.JSON(200, response)
}

func ErrorResponse(c *gin.Context, opts ...ResponseOptions) {
	response := Response{
		Message: "error",
		Code: constant.Fail,
		Data: nil,
	}
	for _, opt := range opts {
		opt(&response)
	}
	c.JSON(500, response)
}