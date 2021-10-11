package server

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type BidInfoRequest struct {
	AuctionId int32 `json:"auctionId" binding:"required"`
	Price     int64 `json:"price" binding:"required"`
}

func (s *Server) DoBid(ctx *gin.Context) {
	var req BidInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseErr(ctx, err, 1)
		return
	}

	userId := s.GetUserId(ctx)
	auction, err := s.store.GetRegisterAuctionByUserId(ctx, db.GetRegisterAuctionByUserIdParams{
		UserID:    userId,
		AuctionID: req.AuctionId,
	})
	if err != nil {
		logrus.Error(err)
		ResponseErrMsg(ctx, err, "User have not permission bid", -1)
		return
	}
	if auction.BidStartDate.After(time.Now()) || time.Now().After(auction.BidEndDate) {
		logrus.Error("Time bid invalid")
		ResponseErrMsg(ctx, err, "Time bid invalid", -1)
		return
	}
	logrus.Infoln(auction)

	if auction.Verify <= 0 {
		logrus.Error("Register auction not verify")
		ResponseErrMsg(ctx, err, "Register auction not verify", -1)
		return
	}

	//check price
	if req.Price < int64(auction.StartPrice) {
		logrus.Error("Price bid invalid ")
		ResponseErrMsg(ctx, err, "Price bid invalid ", -1)
		return
	}

	//check max price
	maxPrice, err := s.store.GetMaxBid(ctx, req.AuctionId)
	if err == nil && maxPrice != nil {
		if maxPrice.(int64) >= req.Price {
			logrus.Infoln("Max price of auction: " + auction.Code + ": " + strconv.FormatInt(maxPrice.(int64), 10))
			logrus.Error("Price bid invalid ")
			ResponseErrMsg(ctx, err, "Price bid invalid ", -1)
			return
		}
	}
	bid, err := s.store.CreateBid(ctx, db.CreateBidParams{
		UserID:    userId,
		AuctionID: req.AuctionId,
		Price:     int32(req.Price),
		Status:    0,
	})
	if err == nil {
		ResponseOK(ctx, bid)
		return
	}
	logrus.Error(err)
	ResponseErr(ctx, err, http.StatusInternalServerError)
}
