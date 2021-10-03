package server

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/config"
	"github.com/bacbia3696/auction/internal/constant"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	store  db.Store
}

func New(cfg *config.Config, store db.Store) *Server {
	return &Server{
		cfg: cfg,
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
	router.POST("/users", server.createUser)
	server.router = router
}
