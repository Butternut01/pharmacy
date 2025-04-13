package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/Butternut01/inventory-service/internal/config"
    "github.com/Butternut01/inventory-service/internal/controller"
    "github.com/Butternut01/inventory-service/internal/repository"
    "github.com/Butternut01/inventory-service/internal/usecase"
)

func main() {
    // Load configuration
    cfg := config.NewConfig()

    // Connect to MongoDB
    db, err := config.ConnectMongoDB(cfg.MongoDBURI)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    // Initialize repository
    productRepo := repository.NewProductRepository(db)

    // Initialize use case
    productUseCase := usecase.NewProductUseCase(productRepo)

    // Initialize controller
    productController := controller.NewProductController(productUseCase)

    // Create Gin router
    router := gin.Default()

    // Define routes
    api := router.Group("/api")
    {
        products := api.Group("/products")
        {
            products.POST("", productController.CreateProduct)
            products.GET("", productController.ListProducts)
            products.GET("/:id", productController.GetProduct)
            products.PATCH("/:id", productController.UpdateProduct)
            products.DELETE("/:id", productController.DeleteProduct)
        }
    }

    // Start server
    if err := router.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}