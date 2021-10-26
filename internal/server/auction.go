package server

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/paycode"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type CreateAuctionRequest struct {
	Title             string                  `form:"title" binding:"required"`
	Description       string                  `form:"description" binding:"required"`
	Code              string                  `form:"code" binding:"required"`
	Owner             string                  `form:"owner" binding:"required,min=6"`
	Organization      string                  `form:"organization" binding:"required"`
	Info              string                  `form:"info" binding:"required"`
	Address           string                  `form:"address" binding:"required"`
	RegisterStartDate time.Time               `form:"registerStartDate" binding:"required"`
	RegisterEndDate   time.Time               `form:"registerEndDate" binding:"required" `
	BidStartDate      time.Time               `form:"bidStartDate" binding:"required"`
	BidEndDate        time.Time               `form:"bidEndDate" binding:"required"`
	StartPrice        int64                   `form:"startPrice" binding:"required"`
	StepPrice         int64                   `form:"stepPrice" binding:"required"`
	Type              int32                   `form:"type" binding:"required"`
	Images            []*multipart.FileHeader `form:"images" binding:"-"`
}
type RespAuctions struct {
	Total    int64         `json:"total"`
	Auctions []AuctionInfo `json:"auctions"`
	PayCode  string        `json:"payCode"`
}
type AuctionInfo struct {
	Auction db.Auction `json:"auction"`
	Images  []string   `json:"images"`
}

func (s *Server) CreateAuction(ctx *gin.Context) {
	res, err := s.createAuction(ctx)
	SendResponse(ctx, res, err)
}

func (s *Server) createAuction(ctx *gin.Context) (interface{}, *ServerError) {
	var req CreateAuctionRequest
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		return nil, ErrUserUnauthorized
	}
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(translateErr(err))
		return nil, ErrInvalidRequest.WithDevMsg(translateErr(err))
	}
	// check code
	_, err := s.store.GetByCode(ctx, req.Code)
	if err == nil {
		return nil, ErrAuctionCodeExisted
	} else if err != sql.ErrNoRows {
		logrus.Error(err)
		return nil, ErrGeneric
	}
	// check date
	if req.BidStartDate.After(req.BidEndDate) || req.RegisterStartDate.After(req.RegisterEndDate) || req.RegisterStartDate.After(req.BidStartDate) || req.RegisterEndDate.After(req.BidEndDate) || req.RegisterEndDate.After(req.BidStartDate) {
		return nil, ErrAuctionDateInvalid
	}
	params := db.CreateAuctionParams{}
	copier.Copy(&params, req)

	auction, err := s.store.CreateAuction(ctx, params)
	if err != nil {
		logrus.Error(err)
		return nil, ErrGeneric
	}
	auctionId := auction.ID

	//handle img
	imgForm, _ := ctx.MultipartForm()
	images := imgForm.File["images"]
	for i := 0; i < len(images); i++ {
		fileName := fmt.Sprintf("static/img/auction_%d_%s", auctionId, RandStringRunes(8)+filepath.Ext(images[i].Filename))
		err = ctx.SaveUploadedFile(images[i], fileName)
		if err != nil {
			logrus.Error(err)
			return nil, ErrGeneric
		}
		_, err = s.store.CreateAuctionImage(ctx, db.CreateAuctionImageParams{
			AuctionID: int64(auctionId),
			Url:       fileName,
		})
		if err != nil {
			logrus.Error(err)
			return nil, ErrGeneric
		}
	}
	return auction, nil
}

func (s *Server) VerifyAuction(ctx *gin.Context) {
	res, err := s.verifyAuction(ctx)
	SendResponse(ctx, res, err)
}

func (s *Server) verifyAuction(ctx *gin.Context) (interface{}, *ServerError) {
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		return nil, ErrUserUnauthorized
	}
	uid, _ := strconv.Atoi(ctx.Query("auctionId"))
	_, err := s.store.GetAuctionById(ctx, int64(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := s.store.UpdateStatusAuction(ctx, db.UpdateStatusAuctionParams{
				Status: 1,
				ID:     int64(uid),
			})
			if err != nil {
				logrus.Error(err)
				return nil, ErrGeneric
			}
		}
		logrus.Error(err)
		return nil, ErrGeneric
	}
	return nil, nil
}

func (s *Server) ListAuction(ctx *gin.Context) {
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

	auctions, err := s.store.GetListAuction(ctx, db.GetListAuctionParams{Code: "%" + keyword + "%", Limit: limit, Offset: offset})
	count, err := s.store.GetTotalAuction(ctx, "%"+keyword+"%")

	var auctionsInfo []AuctionInfo

	for i := 0; i < len(auctions); i++ {
		images, _ := s.store.ListAuctionImage(ctx, auctions[i].ID)
		auction := AuctionInfo{
			Auction: auctions[i],
			Images:  images,
		}
		auctionsInfo = append(auctionsInfo, auction)
	}

	data := RespAuctions{
		Total:    count,
		Auctions: auctionsInfo,
	}
	if err == nil {
		ResponseOK(ctx, data)
		return
	}
	logrus.Error(err)
	ResponseErrMsg(ctx, nil, " Fail", -1)
}

func (s *Server) RegisterAuction(ctx *gin.Context) {
	userId := s.GetUserId(ctx)
	auctionId, _ := strconv.Atoi(ctx.Query("auctionId"))
	auction, err := s.store.GetAuctionById(ctx, int64(auctionId))
	if err == nil {
		if (db.Auction{}) != auction {
			endRegister := auction.RegisterEndDate

			if endRegister.After(time.Now()) {
				ResponseErrMsg(ctx, nil, "Register Auction expired", -1)
				return
			}
			check, err := s.store.GetRegisterAuctionByUserId(ctx, db.GetRegisterAuctionByUserIdParams{
				UserID:    userId,
				AuctionID: int64(auctionId),
			})
			if (db.GetRegisterAuctionByUserIdRow{}) != check {
				ResponseErrMsg(ctx, nil, "User registered", -1)
				return
			}
			res, err := s.store.CreateRegisterAuction(ctx, db.CreateRegisterAuctionParams{
				AuctionID: int64(auctionId),
				UserID:    userId,
				Status:    0,
			})
			if err == nil {
				ResponseOK(ctx, res)
				return
			}
		} else {
			logrus.Error(err)
			ResponseErrMsg(ctx, nil, "Auction not found", -1)
			return
		}
	} else {
		logrus.Error(err)
		ResponseErrMsg(ctx, nil, "Auction not fond", -1)
		return
	}
	ResponseErrMsg(ctx, nil, "Fail", -1)
}

func (s *Server) VerifyRegisterAuction(ctx *gin.Context) {
	roleId := s.GetRoleId(ctx)
	if roleId > 2 {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	uid, _ := strconv.Atoi(ctx.Query("id"))
	check, err := s.store.GetRegisterAuctionById(ctx, int64(uid))
	if err == nil {
		if (db.RegisterAuction{}) != check {
			_, _ = s.store.UpdateStatusRegisterAuction(ctx, db.UpdateStatusRegisterAuctionParams{
				Status: 1,
				ID:     int64(uid),
			})
		}
	}
	ResponseOK(ctx, nil)
}

func (s *Server) ListRegisterAuctionOfUser(ctx *gin.Context) {
	uid := s.GetUserId(ctx)
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
	limit := size
	offset := limit * (page - 1)

	auctions, err := s.store.GetListRegisterAuctionByUserId(ctx, db.GetListRegisterAuctionByUserIdParams{
		UserID: int64(uid),
		Limit:  limit,
		Offset: offset,
	})
	if err == nil {
		ResponseOK(ctx, auctions)
		return
	}
	logrus.Error(err)
	ResponseOK(ctx, nil)
}

func (s *Server) GetAuctionDetail(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("id"))
	auction, err := s.store.GetAuctionById(ctx, int64(aid))
	if err == nil {
		images, _ := s.store.ListAuctionImage(ctx, int64(aid))
		auction := AuctionInfo{
			Auction: auction,
			Images:  images,
		}
		ResponseOK(ctx, auction)
		return
	}
	logrus.Error(err)
	ResponseOK(ctx, nil)
}

func (s *Server) GetAuctionStatus(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("id"))
	uid := s.GetUserId(ctx)
	auction, err := s.store.GetRegisterAuctionByUserId(ctx, db.GetRegisterAuctionByUserIdParams{
		UserID:    int64(uid),
		AuctionID: int64(aid),
	})
	status := -2

	if err == nil {
		status = int(auction.Verify)
	}
	ResponseOK(ctx, status)
}
func (s *Server) GetAuctionPayCode(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("id"))
	uid := s.GetUserId(ctx)
	user, err := s.store.GetById(ctx, uid)
	if err == nil {
		payCode := paycode.NewFromAuctionID(user.UserName, aid)
		ResponseOK(ctx, payCode)
		return
	}
	ResponseOK(ctx, nil)
}

func (s *Server) GetMaxBidAuction(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("id"))
	uid := s.GetUserId(ctx)
	auction, err := s.store.GetRegisterAuctionByUserId(ctx, db.GetRegisterAuctionByUserIdParams{
		UserID:    uid,
		AuctionID: int64(aid),
	})
	if err != nil || auction.Verify <= 0 {
		logrus.Error(err)
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	maxPrice, err := s.store.GetMaxBid(ctx, int64(aid))
	if err == nil && maxPrice != nil {
		ResponseOK(ctx, maxPrice)
		return
	}
	ResponseOK(ctx, auction.StartPrice)
}

func (s *Server) GetWinnerAuction(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("id"))
	uid := s.GetUserId(ctx)
	auction, err := s.store.GetRegisterAuctionByUserId(ctx, db.GetRegisterAuctionByUserIdParams{
		UserID:    uid,
		AuctionID: int64(aid),
	})
	if err != nil || auction.Verify <= 0 {
		logrus.Error(err)
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	winner, err := s.store.GetWinnerAuction(ctx, int64(aid))
	if err == nil {
		ResponseOK(ctx, winner)
		return
	}
	ResponseOK(ctx, nil)
}

func (s *Server) checkPermission(ctx *gin.Context, uid, aid int64) bool {
	auction, err := s.store.GetRegisterAuctionByUserId(ctx, db.GetRegisterAuctionByUserIdParams{
		UserID:    int64(uid),
		AuctionID: int64(aid),
	})
	if err != nil || auction.Verify <= 0 {
		logrus.Error(err)
		return false
	}
	return true
}

type RespUsersRegisterAuction struct {
	Total int64                                 `json:"total"`
	Users []db.GetAllListUserRegisterAuctionRow `json:"users"`
}
type RespUsersRegisterAuctionByStatus struct {
	Total int64                                      `json:"total"`
	Users []db.GetListUserRegisterAuctionByStatusRow `json:"users"`
}

func (s *Server) GetListUserRegisterAuction(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseErr(ctx, err, 1)
		return
	}
	uid := s.GetUserId(ctx)
	status := 1
	flag := true
	roleId := s.GetRoleId(ctx)
	if roleId < 3 {
		status = req.Status
	} else {
		flag = s.checkPermission(ctx, uid, req.AuctionId)
	}
	if flag == false {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	size := req.Size
	if size == 0 {
		size = 1000
	}
	limit := size
	offset := limit * (page - 1)
	if status == 2 {
		users, err := s.store.GetAllListUserRegisterAuction(ctx, db.GetAllListUserRegisterAuctionParams{
			AuctionID: req.AuctionId,
			Offset:    offset,
			Limit:     limit,
		})
		count, err := s.store.GetTotalUserRegisterAuction(ctx, req.AuctionId)

		if err == nil {
			resp := RespUsersRegisterAuction{
				Users: users,
				Total: count,
			}
			ResponseOK(ctx, resp)
			return
		}
		logrus.Error(err)
	} else {
		users, err := s.store.GetListUserRegisterAuctionByStatus(ctx, db.GetListUserRegisterAuctionByStatusParams{
			AuctionID: req.AuctionId,
			Offset:    offset,
			Limit:     limit,
			Status:    int32(status),
		})
		count, err := s.store.GetTotalUserRegisterAuctionByStatus(ctx, db.GetTotalUserRegisterAuctionByStatusParams{
			AuctionID: req.AuctionId,
			Status:    int32(status),
		})

		if err == nil {
			resp := RespUsersRegisterAuctionByStatus{
				Users: users,
				Total: count,
			}
			ResponseOK(ctx, resp)
			return
		}
		logrus.Error(err)
	}
	ResponseOK(ctx, nil)
}

type RespUsersBidAuction struct {
	Total int64                            `json:"total"`
	Users []db.GetAllListUserBidAuctionRow `json:"users"`
}

func (s *Server) GetListUserBiAuction(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseErr(ctx, err, 1)
		return
	}
	auction, _ := s.store.GetAuctionById(ctx, req.AuctionId)
	if auction.BidEndDate.After(time.Now()) {
		logrus.Info("Auction have not end")
		ResponseOK(ctx, nil)
		return
	}
	uid := s.GetUserId(ctx)
	flag := true
	roleId := s.GetRoleId(ctx)
	if roleId >= 3 {
		flag = s.checkPermission(ctx, int64(uid), req.AuctionId)
	}
	if flag == false {
		ResponseErrMsg(ctx, nil, "User have not permission", -1)
		return
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	size := req.Size
	if size == 0 {
		size = 1000
	}
	limit := size
	offset := limit * (page - 1)
	users, err := s.store.GetAllListUserBidAuction(ctx, db.GetAllListUserBidAuctionParams{
		AuctionID: req.AuctionId,
		Offset:    offset,
		Limit:     limit,
	})
	count, err := s.store.GetTotalListUserBidAuction(ctx, req.AuctionId)

	if err == nil {
		resp := RespUsersBidAuction{
			Users: users,
			Total: count,
		}
		ResponseOK(ctx, resp)
		return
	}
	logrus.Error(err)
	ResponseOK(ctx, nil)
}
