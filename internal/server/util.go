package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strings"
	"text/template"
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
		_ = entranslations.RegisterDefaultTranslations(v, trans) //nolint:errcheck
		return
	}
	logrus.Info("Error trying to setup translation")
}

func ResponseOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": 0,
		"msg":  "success",
	})
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

func ResponseErrMsg(ctx *gin.Context, data interface{}, msg string, code int) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": code,
		"msg":  msg,
	})
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func Format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	return b.String()
}