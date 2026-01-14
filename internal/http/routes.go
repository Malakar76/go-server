package http

import (
	"go-server/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, h *handlers.Handlers) {
	r.GET("/health", h.Health)

	v1 := r.Group("/v1")
	{
		v1.POST("/users", h.CreateUser)
		v1.GET("/users/:id", h.GetUser)
	}
}
