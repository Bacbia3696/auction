package server

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func (s *Server) GetRoleId(ctx *gin.Context) int32 {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Claims)
	userId := authPayload.Id
	roleId, _ := s.store.GetRoleByUserId(ctx, userId)
	return roleId
}

func (s *Server) ListUser(ctx *gin.Context) {
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	users, err := s.store.GetListUser(ctx)
	if err == nil {
		ResponseOK(ctx, users)
		return
	}
	ResponseErrMsg(ctx, nil, " Fail", -1)
}

func (s *Server) VerifyUser(ctx *gin.Context) {
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	uid, _ := strconv.Atoi(ctx.Query("userId"))
	checkUser, err := s.store.GetById(ctx, int32(uid))
	if err == nil {
		if (db.User{}) != checkUser {
			_, _ = s.store.UpdateStatus(ctx, db.UpdateStatusParams{
				Status: 1,
				ID:     int32(uid),
			})
		}
	}
	ResponseOK(ctx, nil)
}

func (s *Server) LockUser(ctx *gin.Context) {
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	uid, _ := strconv.Atoi(ctx.Query("userId"))
	checkUser, err := s.store.GetById(ctx, int32(uid))
	if err == nil {
		if (db.User{}) != checkUser {
			_, _ = s.store.UpdateStatus(ctx, db.UpdateStatusParams{
				Status: -1,
				ID:     int32(uid),
			})
		}
	}
	ResponseOK(ctx, nil)
}
