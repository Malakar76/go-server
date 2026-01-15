package http

import (
	"go-server/internal/config"
	"go-server/internal/http/handlers"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	addr   string
}

func NewServer(cfg config.Config, h *handlers.Handlers) *Server {
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(gin.Logger(), gin.Recovery())
	registerRoutes(r, h)

	return &Server{engine: r, addr: cfg.HTTPAddr}
}

func (s *Server) Run() error {
	return s.engine.Run(s.addr)
}
