package main

import (
	"context"
	"log"

	"github.com/AhmadMuj/books-api-go/internal/cache"
	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/AhmadMuj/books-api-go/internal/events"
	"github.com/AhmadMuj/books-api-go/internal/handlers"
	"github.com/AhmadMuj/books-api-go/internal/repository"
	"github.com/AhmadMuj/books-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

// @title Books API Go
// @version 1.0
// @description A RESTful API for managing books with Kafka event streaming and Redis caching
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, _ := config.LoadConfig(".env")

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	cacheInstance, err := cache.NewRedisCache(cfg)

	if err != nil {
		log.Fatal("Failed to initialize cache:", err)
	}
	defer cacheInstance.Close()

	kafkaProducer, err := events.NewKafkaProducer(cfg)
	if err != nil {
		log.Fatal("Failed to initialize Kafka producer:", err)
	}
	defer kafkaProducer.Close()

	// Initialize event service
	eventService := events.NewEventService(kafkaProducer)

	// Optionally initialize consumer
	kafkaConsumer, err := events.NewKafkaConsumer(cfg)
	if err != nil {
		log.Fatal("Failed to initialize Kafka consumer:", err)
	}
	defer kafkaConsumer.Close()

	if err := kafkaConsumer.Start(context.Background()); err != nil {
		log.Fatal("Failed to start Kafka consumer:", err)
	} else {
		log.Println("Kafka consumer started")
	}

	// Initialize repository
	bookRepo := repository.NewBookRepository(db.DB)

	// Initialize service
	bookService := service.NewBookService(bookRepo, cacheInstance, eventService)

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
