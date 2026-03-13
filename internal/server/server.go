package server

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	UserHandler(router *gin.Engine)
}

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {

	return &Server{router: router}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
