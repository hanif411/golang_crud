package controllers

import (
	"golang_crud/database"
	"golang_crud/models"
	"golang_crud/services"
	"golang_crud/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	svc *services.ProductService
}

func NewProductController(svc *services.ProductService) *ProductController {
	return &ProductController{svc: svc}
}

func CreateProduct(c *gin.Context) {

	var product models.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "Invalid product", Data: nil})
		return
	}

	key := []byte(os.Getenv("KEY_CRYPTO"))
	product.SecretNote = utils.Encrypt((key), product.SecretNote)

	result := database.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error create ke database", Data: nil})
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusCreated, Message: "Sukses create product", Data: product})
}

func (ctrl *ProductController) GetAllProduct(c *gin.Context) {
	key := []byte(os.Getenv("KEY_CRYPTO"))
	products, err := ctrl.svc.FindAll(key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Success get All product", Data: products})
}

func GetProductById(c *gin.Context) {

	id := c.Param("id")
	var product models.Product

	result := database.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.BaseResponse{Status: http.StatusNotFound, Message: "Product not found", Data: nil})
		return
	}

	product.SecretNote = utils.Decrypt([]byte(os.Getenv("KEY_CRYPTO")), product.SecretNote)

	c.JSON(http.StatusOK, models.BaseResponse{
		Status: http.StatusOK, Message: "Success Get Product", Data: product,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "invalid id", Data: nil})
	}
	var product models.Product
	err = c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "Bad request product", Data: nil})
		return
	}

	product.ID = uint(idNumber)
	key := []byte(os.Getenv("KEY_CRYPTO"))
	product.SecretNote = utils.Encrypt(key, product.SecretNote)

	result := database.DB.Model(&product).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "Failed create to database", Data: nil})
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{
		Status: http.StatusOK, Message: "sukses update", Data: product,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "Failed to create to database", Data: nil})
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{
		Status: http.StatusOK, Message: "Success Delete product", Data: nil,
	})
}
