package handler

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	UserService UserService
}

type Handler struct {
	userHandler *UserHandler
}

func NewHandler(router *gin.Engine, svc Service) *Handler {
	userHandler := NewUserHandler(svc.UserService)

	// Register routes
	router.GET("/users", userHandler.GetUsers)

	return &Handler{userHandler: userHandler}
}
