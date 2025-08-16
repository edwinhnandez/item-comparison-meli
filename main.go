package main

import (
	"log"

	"item-comparison-meli.com/m/api/controllers"
	"item-comparison-meli.com/m/api/services"
	_ "item-comparison-meli.com/m/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Item Comparison API
// @version 1.0
// @description API for comparing product items
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize product service
	productService, err := services.NewProductService("data/products.json")
	if err != nil {
		log.Fatalf("Failed to initialize product service: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Initialize controllers
	productController := controllers.NewProductController(productService)

	// API routes
	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", productController.GetProducts)
			products.GET("/compare", productController.CompareProducts)
			products.GET("/:id", productController.GetProduct)
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
