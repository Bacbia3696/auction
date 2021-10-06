package server

import (
	"fmt"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) ListUser(ctx *gin.Context) {
	var req UserLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user, err := s.store.GetByUserName(ctx, req.UserName)
		if err == nil {
			if CheckPasswordHash(req.Password, user.Password) {
				token, _ := token.GenToken(user)
				ResponseOK(ctx, token)
			} else {
				ResponseErrMsg(ctx, nil, "Unauthorized", 401)
			}
		} else {
			ResponseErrMsg(ctx, nil, "Unauthorized", 401)
		}
	}
}

func (s *Server) VerifyUser(ctx *gin.Context) {
	userId := ctx.Query("userId")
	fmt.Print(userId)
	ResponseErrMsg(ctx, nil, "okie", 401)

}

func (s *Server) LockUser(ctx *gin.Context) {
	var req UserLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user, err := s.store.GetByUserName(ctx, req.UserName)
		if err == nil {
			if CheckPasswordHash(req.Password, user.Password) {
				token, _ := token.GenToken(user)
				ResponseOK(ctx, token)
			} else {
				ResponseErrMsg(ctx, nil, "Unauthorized", 401)
			}
		} else {
			ResponseErrMsg(ctx, nil, "Unauthorized", 401)
		}
	}
}
