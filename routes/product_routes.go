package routes

import (
	"golang_crud/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, ctrl *controllers.ProductController) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/", ctrl.GetAllProduct)
	}
}
