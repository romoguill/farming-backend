package server

import (
	"github.com/gin-gonic/gin"
	"github.com/romoguill/farming-backend/internal/handler"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	server := &Server{router: router}

	router.GET("/health", handler.HealthCheck)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
