package server

import (
	"github.com/gin-gonic/gin"
	"github.com/romoguill/farming-backend/internal/service"
)

type Handler interface {
	RegisterRoutes(router *gin.Engine)
}

type Server struct {
	router  *gin.Engine
	handler Handler
	service *service.Service
}

func NewServer(handler Handler, service service.Service) *Server {
	router := gin.Default()
	h := handler.NewHandler(router)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
