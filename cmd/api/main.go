package main

import (
	"log"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	gin.SetMode(cfg.Server.Mode)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
