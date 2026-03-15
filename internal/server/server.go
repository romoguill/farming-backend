package server

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	GetUsers(*gin.Context)
}

type server struct {
	handlers Handlers
	router   *gin.Engine
}

func NewServer(handlers Handlers) *server {
	return &server{handlers: handlers}
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
