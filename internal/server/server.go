package server

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/config"
	"github.com/bacbia3696/auction/internal/constant"
	"github.com/bacbia3696/auction/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
	// TODO: bind custom validator
	s.setupRouter()
	transInit()
	return s.router.Run(buildAddr(s.cfg.Server.Host, s.cfg.Server.Port))
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/static", static.LocalFile("./static", false)))
	v1 := router.Group("/user")
	{
		v1.POST("/register", server.RegisterUser)
		v1.POST("/login", server.LoginUser)
		v1.POST("/auctions", server.ListAuction)

	}

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/cms/verify", server.VerifyUser)
		authRoutes.GET("/cms/lock", server.LockUser)
		authRoutes.POST("/cms/list-user", server.ListUser)

		authRoutes.POST("/cms/auction/create", server.CreateAuction)
		authRoutes.GET("/cms/auction/verify", server.VerifyAuction)
		authRoutes.GET("/cms/register-auction/verify", server.VerifyRegisterAuction)

		authRoutes.POST("/user/change-password", server.ChangePassword)
		authRoutes.GET("/user/register-auction", server.RegisterAuction)
		authRoutes.POST("/user/list/register-auction", server.ListRegisterAuction)

	}
	server.router = router
}
