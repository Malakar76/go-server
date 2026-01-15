package handlers

import (
	"net/http"

	"go-server/internal/repository"

	"github.com/gin-gonic/gin"
)

type CreateObjectRequest struct {
	Name   string `json:"name" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func (h *Handlers) CreateObject(c *gin.Context) {
	var req CreateObjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o, err := h.objectSvc.Create(c.Request.Context(), req.Name, req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, o)
}

// GetObject godoc
// @Summary      Get an object by ID
// @Tags         objects
// @Produce      json
// @Param        id   path      string  true  "Object ID"
// @Success      200  {object}  domain.Object
// @Failure      404  {object}  map[string]string
// @Router       /v1/objects/{id} [get]
func (h *Handlers) GetObject(c *gin.Context) {
	o, err := h.objectSvc.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		if err == repository.ErrObjectNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "object not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, o)
}
