package server

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/config"
	"github.com/bacbia3696/auction/internal/constant"
	"github.com/bacbia3696/auction/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	store  db.Store
}

func New(cfg *config.Config, store db.Store) *Server {
	return &Server{
		cfg:   cfg,
		store: store,
	}
}

func (s *Server) Serve() error {
	if s.cfg.Environment == constant.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}
	if s.cfg.LogFile != "" {
		gin.DisableConsoleColor()
	}
	gin.DefaultWriter = logrus.StandardLogger().Out
	// TODO: bind custom validator
	s.setupRouter()
	transInit()
	return s.router.Run(buildAddr(s.cfg.Server.Host, s.cfg.Server.Port))
}

func (server *Server) setupRouter() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))
	router.Use(static.Serve("/static", static.LocalFile("./static", false)))
	router.GET("/ws", server.wsHandlerFunc)
	v1 := router.Group("/user")
	{
		v1.POST("/register", server.RegisterUser)
		v1.POST("/login", server.LoginUser)
		v1.POST("/auctions", server.ListAuction)
	}
	v2 := router.Group("/auction")
	{
		v2.GET("/detail", server.GetAuctionDetail)
	}

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/cms/verify", server.VerifyUser)
		authRoutes.POST("/cms/lock", server.LockUser)
		authRoutes.POST("/cms/list-user", server.ListUser)
		authRoutes.GET("/cms/user/detail", server.ViewDetailUser)

		authRoutes.POST("/cms/auction/create", server.CreateAuction)
		authRoutes.POST("/cms/auction/verify", server.VerifyAuction)
		authRoutes.POST("/cms/register-auction/verify", server.VerifyRegisterAuction)

		authRoutes.POST("/user/change-password", server.ChangePassword)
		authRoutes.POST("/user/register-auction", server.RegisterAuction)
		authRoutes.POST("/user/list/register-auction", server.ListRegisterAuctionOfUser)
		authRoutes.POST("/user/auction/bid", server.DoBid)
		authRoutes.GET("/user/info", server.GetUserInfo)
		authRoutes.GET("/user/auction/status", server.GetAuctionStatus)
		authRoutes.GET("/auction/max-price", server.GetMaxBidAuction)
		authRoutes.POST("/auction/user-register", server.GetListUserRegisterAuction)
		authRoutes.GET("/auction/pay-code", server.GetAuctionPayCode)
		authRoutes.POST("/auction/user-bid", server.GetListUserBiAuction)
		authRoutes.GET("/auction/winner", server.GetWinnerAuction)

	}
	server.router = router
}
