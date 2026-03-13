package main

import (
	"github.com/romoguill/farming-backend/internal/server"
	"github.com/romoguill/farming-backend/internal/service"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository.NewRepository())

	address := "localhost:3000"
	server.Start(address)
}
