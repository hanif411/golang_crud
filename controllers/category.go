package controllers

import (
	"golang_crud/database"
	"golang_crud/models"
	"golang_crud/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

	var category models.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "Invalid category", Data: nil})
		return
	}

	result := database.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "Invalid create category to database" + result.Error.Error(), Data: nil})
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Succes create category", Data: category})
}

func GetAllCategory(c *gin.Context) {
	var category models.Category
	err := database.DB.Find(&category)
	if err != nil {
		c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Succes get all category", Data: category})
	}
}

func GetCategoryByIdFindProduct(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	result := database.DB.Preload("Products").First(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.BaseResponse{Status: http.StatusNotFound, Message: "Category not found", Data: nil})
		return
	}

	key := []byte(os.Getenv("KEY_CRYPTO"))
	for i := range category.Products {
		category.Products[i].SecretNote = utils.Decrypt(key, category.Products[i].SecretNote)
	}
	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Success get category by id and product", Data: category})
}

