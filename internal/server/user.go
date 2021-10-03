package server

import (
	"net/http"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type createUserRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName" binding:"required"`
	// Email    string `json:"email" binding:"required,email"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responeErr(ctx, err, 1)
		return
	}
	// create new user
	params := db.CreateUserParams{}
	copier.Copy(&params, req)
	logrus.Debug("params", params)
	user, err := s.store.CreateUser(ctx, params)
	if err != nil {
		responseOK(ctx, user)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
