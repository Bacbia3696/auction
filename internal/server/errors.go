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

func (e *ServerError) DevMsg() string {
	return e.devMsg
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
		"code":   err.Code(),
		"msg":    err.Error(),
		"devMsg": err.DevMsg(),
		"data":   nil,
	})
}

var (
	ErrGeneric             = NewError(10_000, "Có lỗi xảy ra, vui lòng thử lại sau!")
	ErrInvalidLogin        = NewError(10_001, "Thông tin đăng nhập không đúng, vui lòng thử lại!")
	ErrInvalidRequest      = NewError(10_002, "Request không hợp lệ!")
	ErrUserBlock           = NewError(10_003, "Tài khoản đã bị khoá, vui lòng liên hệ tổng đài để được hỗ trợ!")
	ErrUserUnauthorized    = NewError(10_004, "Tài khoản không có quyền truy cập nội dung này!")
	ErrUserNameExisted     = NewError(10_005, "Username đã tồn tại, vui lòng chọn username khác!")
	ErrEmailExisted        = NewError(10_006, "Email đã tồn tại, vui lòng chọn email khác!")
	ErrAuctionCodeExisted  = NewError(10_007, "Mã đấu giá đã tồn tại!")
	ErrAuctionDateInvalid1 = NewError(10_008, "Ngày kết thúc phải lớn hơn ngày bắt đầu!")
	ErrAuctionDateInvalid2 = NewError(10_009, "Ngày bắt đầu đấu giá phải lớn hơn ngày kết thúc đăng ký!")

	ErrNotCreateWsConnection = NewError(20_000, "Không thể kết nối, vui lòng thử lại!")
)
