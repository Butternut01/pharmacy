package main

import (
    "log"
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/Butternut01/api-gateway/internal/config"
    "github.com/Butternut01/api-gateway/internal/handler"
    "github.com/Butternut01/api-gateway/internal/middleware"
)

func main() {
    // Load configuration
    cfg := config.NewConfig()

    // Create Gin router
    router := gin.Default()

    // Middleware
    router.Use(middleware.LoggingMiddleware())
    router.Use(middleware.AuthMiddleware())

    // Routes
    inventoryProxy := handler.NewInventoryServiceProxy(cfg.InventoryServiceURL)
    orderProxy := handler.NewOrderServiceProxy(cfg.OrderServiceURL)

    router.Any("/inventory/*path", inventoryProxy)
    router.Any("/orders/*path", orderProxy)

    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // Start server
    if err := router.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}