package main

import (
	"github.com/romoguill/farming-backend/internal/server"
)

func main() {
	server := server.NewServer()

	address := "localhost:3000"
	server.Start(address)
}
