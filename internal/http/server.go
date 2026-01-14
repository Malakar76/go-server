package http

import (
	"go-server/internal/config"
	"go-server/internal/http/handlers"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	addr   string
}

func NewServer(cfg config.Config, userSvc *service.UserService) *Server {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.New(userSvc)
	registerRoutes(r, h)

	return &Server{engine: r, addr: cfg.HTTPAddr}
}

func (s *Server) Run() error {
	return s.engine.Run(s.addr)
}
