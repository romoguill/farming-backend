package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romoguill/farming-backend/internal/model"
)

type UserService interface {
	GetUsers() ([]model.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}
