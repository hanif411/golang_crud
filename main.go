package main

import (
	"fmt"
	"golang_crud/controllers"
	"golang_crud/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")

	}

	database.ConnectDatabase()

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		product := api.Group("/product")
		{
			product.GET("/", controllers.GetAllProduct)
			product.POST("/", controllers.CreateProduct)
			product.GET("/:id", controllers.GetProductById)
			product.PUT("/:id", controllers.UpdateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)
		}
		category := api.Group("/category")
		{
			category.POST("/", controllers.CreateCategory)
			category.GET("/", controllers.GetAllCategory)
			category.GET("/:id", controllers.GetCategoryByIdFindProduct)
		}
	}

	fmt.Println("server jalan di port 5000")
	r.Run(":5000")
}
