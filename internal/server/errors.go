package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerError struct {
	code int
	msg  string
}

func NewError(code int, msg string) *ServerError {
	return &ServerError{
		code: code,
		msg:  msg,
	}
}

func (e *ServerError) Error() string {
	return e.msg
}

func (e *ServerError) Code() int {
	return e.code
}

func (e *ServerError) WithCustomMessage(msg string) *ServerError {
	return &ServerError{
		code: e.code,
		msg:  msg,
	}
}

func SendResponse(ctx *gin.Context, res interface{}, err *ServerError) {
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err.Code(),
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": res,
		})
	}
}

var (
	ErrGeneric          = NewError(10_000, "Có lỗi xảy ra, vui lòng thử lại sau")
	ErrInvalidLogin     = NewError(10_001, "Thông tin đăng nhập không đúng, vui lòng thử lại")
	ErrInvalidRequest   = NewError(10_002, "Request không hợp lệ")
	ErrUserBlock        = NewError(10_003, "Tài khoản đã bị khoá, vui lòng liên hệ tổng đài để được hỗ trợ")
	ErrUserUnauthorized = NewError(10_004, "Tài khoản không có quyền truy cập nội dung này")
)
