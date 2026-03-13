package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	router *gin.Engine
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", HealthCheck)
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{router: router}
}
