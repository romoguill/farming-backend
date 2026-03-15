package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romoguill/farming-backend/internal/model"
)

type UserService interface {
	GetAll() ([]model.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserDTO struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUsersResponse []UserDTO

func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := handler.userService.GetAll()
	if err != nil {
		log.Printf("error in get users handler: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error getting users"})
	}

	var res GetUsersResponse
	for _, user := range users {
		res = append(res, UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	ctx.BindJSON(res)
}
