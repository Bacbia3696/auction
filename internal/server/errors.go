package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerError struct {
	code   int
	msg    string
	devMsg string
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

func (e *ServerError) WithDevMsg(m string) *ServerError {
	return &ServerError{
		code:   e.code,
		msg:    e.msg,
		devMsg: m,
	}
}

func SendResponse(ctx *gin.Context, res interface{}, err *ServerError) {
	if err != nil {
		SendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
	})
}

func SendErr(ctx *gin.Context, err *ServerError) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": err.Code(),
		"msg":  err.Error(),
		"data": nil,
	})
}

var (
	ErrGeneric          = NewError(10_000, "Có lỗi xảy ra, vui lòng thử lại sau")
	ErrInvalidLogin     = NewError(10_001, "Thông tin đăng nhập không đúng, vui lòng thử lại")
	ErrInvalidRequest   = NewError(10_002, "Request không hợp lệ")
	ErrUserBlock        = NewError(10_003, "Tài khoản đã bị khoá, vui lòng liên hệ tổng đài để được hỗ trợ")
	ErrUserUnauthorized = NewError(10_004, "Tài khoản không có quyền truy cập nội dung này")

	ErrNotCreateWsConnection = NewError(20_000, "Không thể kết nối, vui lòng thử lại!")
)
