package main

import (
	"inventory-service/infrastructure/db"
	"inventory-service/internal/handler"
	"inventory-service/internal/repository"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация MongoDB
	client, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	defer client.Disconnect(nil)

	categoryRepo := repository.NewCategoryRepository(client.Database("pharmacy"))
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	
	

	// Создаем репозиторий для работы с продуктами
	productRepo := repository.NewProductRepository(client.Database("pharmacy"))

	// Создаем обработчики для продуктов
	productHandler := handler.NewProductHandler(productRepo)

	// Инициализируем Gin
	r := gin.Default()

	// Маршруты для работы с продуктами
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.PATCH("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)
	r.GET("/products", productHandler.GetAllProducts)

	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories", categoryHandler.GetAllCategories)
	
	r.Run(":8081") // inventory-service будет на этом порту
}
