package main

import (
    "log"

    "github.com/Butternut01/order-service/internal/config"
    "github.com/Butternut01/order-service/internal/controller"
    "github.com/Butternut01/order-service/internal/repository"
    "github.com/Butternut01/order-service/internal/usecase"
    "github.com/gin-gonic/gin"
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
	orderRepo := repository.NewOrderRepository(db)

	// Initialize use case
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Initialize controller
	orderController := controller.NewOrderController(orderUseCase)

	// Create Gin router
	router := gin.Default()

	// Define routes
	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.POST("", orderController.CreateOrder)
			orders.GET("", orderController.ListOrders)
			orders.GET("/:id", orderController.GetOrder)
			orders.PATCH("/:id", orderController.UpdateOrderStatus)
		}
	}

	// Start server
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
