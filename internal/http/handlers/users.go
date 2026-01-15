package handlers

import (
	"net/http"

	"go-server/internal/repository"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handlers) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userSvc.Create(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

// GetUser godoc
// @Summary      Get a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  domain.User
// @Failure      404  {object}  map[string]string
// @Router       /v1/users/{id} [get]
func (h *Handlers) GetUser(c *gin.Context) {
	u, err := h.userSvc.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, u)
}
