package server

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func (s *Server) GetUserId(ctx *gin.Context) int32 {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Claims)
	userId := authPayload.Id
	return userId
}

type Request struct {
	Keyword string `json:"keyword"`
	Page    int32  `json:"page"`
	Size    int32  `json:"size"`
}

type RespUsers struct {
	Total int64 `json:"total"`
	Users    []db.User  `json:"users"`
}

func (s *Server) ListUser(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseErr(ctx, err, 1)
		return
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	size := req.Size
	if size == 0 {
		size = 10
	}
	keyword := req.Keyword
	limit := size
	offset := limit * (page - 1)

	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	users, err := s.store.GetListUser(ctx, db.GetListUserParams{UserName: "%" + keyword + "%", Limit: limit, Offset: offset})
	count, err := s.store.GetTotalUser(ctx,  "%" + keyword + "%")

	data := RespUsers {
		Total: count,
		Users: users,
	}
	if err == nil {
		ResponseOK(ctx, data)
		return
	}
	logrus.Error(err)
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
