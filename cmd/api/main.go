package main

import (
	"log"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/AhmadMuj/books-api-go/internal/handlers"
	"github.com/AhmadMuj/books-api-go/internal/repository"
	"github.com/AhmadMuj/books-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

// @title Books API Go
// @version 1.0
// @description A RESTful API for managing books with Kafka event streaming and Redis caching
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize repository
	bookRepo := repository.NewBookRepository(db.DB)

	// Initialize service
	bookService := service.NewBookService(bookRepo)

	// Initialize handler
	bookHandler := handlers.NewBookHandler(bookService)

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	handlers.SetupRoutes(r, bookHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
