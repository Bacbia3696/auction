package server

import (
	"net/http"
	"strconv"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/constant"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func (s *Server) GetRoleId(ctx *gin.Context) int64 {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Claims)
	userId := authPayload.Id
	roleId, _ := s.store.GetRoleByUserId(ctx, userId)
	return roleId
}

func (s *Server) GetUserId(ctx *gin.Context) int64 {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Claims)
	userId := authPayload.Id
	return userId
}

type Request struct {
	Keyword   string `json:"keyword"`
	Status    *int   `json:"status"`
	Page      int32  `json:"page"`
	AuctionId int64  `json:"auctionId"`
	Size      int32  `json:"size"`
}

type RespUsers struct {
	Total int64     `json:"total"`
	Users []db.User `json:"users"`
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

	var users []db.User
	var count int64
	var err error
	if req.Status == nil {
		logrus.Println("there")
		users, err = s.store.GetListUser(ctx, db.GetListUserParams{UserName: "%" + keyword + "%", Limit: limit, Offset: offset})
		count, err = s.store.GetTotalUser(ctx, "%"+keyword+"%")
	} else {
		logrus.Println("here")
		users, err = s.store.GetListUserStatusInit(ctx, db.GetListUserStatusInitParams{UserName: "%" + keyword + "%", Limit: limit, Offset: offset})
		count, err = s.store.GetTotalUserStatusInit(ctx, "%"+keyword+"%")
	}

	data := RespUsers{
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
	checkUser, err := s.store.GetById(ctx, int64(uid))
	if err == nil {
		if (db.User{}) != checkUser {
			_, _ = s.store.UpdateStatus(ctx, db.UpdateStatusParams{
				Status: constant.USER_STATUS_VERIFIED,
				ID:     int64(uid),
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
	checkUser, err := s.store.GetById(ctx, int64(uid))
	if err == nil {
		if (db.User{}) != checkUser {
			_, _ = s.store.UpdateStatus(ctx, db.UpdateStatusParams{
				Status: constant.USER_STATUS_LOCKED,
				ID:     int64(uid),
			})
		}
	}
	ResponseOK(ctx, nil)
}

type UserInfo struct {
	User   db.User        `json:"user"`
	Images []db.UserImage `json:"images"`
	RoleId int64          `json:"roleId"`
}

func (s *Server) ViewDetailUser(ctx *gin.Context) {
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	uid, _ := strconv.Atoi(ctx.Query("id"))
	user, err := s.store.GetById(ctx, int64(uid))
	if err == nil {
		images, _ := s.store.ListImage(ctx, int64(uid))
		roleId, _ := s.store.GetRoleByUserId(ctx, int64(uid))
		userInfo := UserInfo{
			User:   user,
			Images: images,
			RoleId: roleId,
		}
		ResponseOK(ctx, userInfo)
		return
	}
	logrus.Error("ViewDetailUser ", err)
	ResponseErr(ctx, err, http.StatusInternalServerError)
	return
}
