package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/romoguill/farming-backend/internal/database"
	"github.com/romoguill/farming-backend/internal/handler"
	"github.com/romoguill/farming-backend/internal/repository"
	"github.com/romoguill/farming-backend/internal/server"
	"github.com/romoguill/farming-backend/internal/service"
)

func main() {

	// Initialize database. App will exit if database connection fails.
	db, err := database.NewDatabase(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Cannot start application. Failed to connect to database: %v.", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db.DB)
	svc := service.NewService(repo.UserRepository)

	router := gin.Default()
	handlers := handler.NewHandler(router, svc)

	server := server.NewServer(handlers)

	address := "localhost:3000"
	server.Start(address)
}
