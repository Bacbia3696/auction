package middleware

import (
	"errors"
	"fmt"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)
const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"

)
var (
	trans ut.Translator
)
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ResponseErr(ctx, err, http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ResponseErr(ctx, err, http.StatusUnauthorized)
			return
		}

		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ResponseErr(ctx, err, http.StatusUnauthorized)
			return
		}

		accessToken := fields[1]
		payload, err := token.Verify(accessToken)
		if err != nil {
			ResponseErrMsg(ctx, nil, "Token invalid", http.StatusUnauthorized)
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
func ResponseErr(ctx *gin.Context, err error, code int) {
	var req gin.H
	msg := err.Error()
	if validatorErrs, ok := err.(validator.ValidationErrors); ok && len(validatorErrs) > 0 {
		msg = ""
		for _, e := range validatorErrs {
			msg += e.Translate(trans) + ", "
		}
		msg = msg[:len(msg)-2]
	}
	req = gin.H{
		"code": code,
		"data": nil,
		"msg":  msg,
	}
	ctx.JSON(http.StatusOK, req)
}

func ResponseErrMsg(ctx *gin.Context, data interface{}, msg string , code int) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": code,
		"msg":  msg,
	})
}
