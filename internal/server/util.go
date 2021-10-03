package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
)

// use a single instance , it caches struct info
var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func buildAddr(host string, port uint32) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func transInit() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni = ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		_ = en_translations.RegisterDefaultTranslations(v, trans)
	}
	logrus.Info("Error trying to setup translation")
}

func responseOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": 0,
		"msg":  "success",
	})
}

func responeErr(ctx *gin.Context, err error, code int) {
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
