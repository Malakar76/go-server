package http

import (
	"go-server/internal/http/handlers"

	_ "go-server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func registerRoutes(r *gin.Engine, h *handlers.Handlers) {
	r.GET("/health", h.Health)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1")
	{
		v1.POST("/users", h.CreateUser)
		v1.GET("/users/:id", h.GetUser)
		v1.GET("/objects/:id", h.GetObject)
		v1.POST("/objects", h.CreateObject)
	}
}
