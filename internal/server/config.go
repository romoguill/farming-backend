package server

import (
	"github.com/gin-gonic/gin"
	"github.com/romoguill/farming-backend/internal/handler"
	"github.com/romoguill/farming-backend/internal/service"
)

type Config struct {
	Port       string
	Handler    *handler.BaseHandler
	Service    *service.Service
}

func NewConfig(router *gin.Engine) *Config {
	return &Config{
		Port:    "8080",
		Handler: handler.NewHandler(router),
		Service: service.NewService(),
	}
}
