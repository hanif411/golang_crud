package database

import (
	"fmt"
	"golang_crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "root:kaoskaki@tcp(127.0.0.1:3306)/go_secure_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database!")
	}

	err = database.AutoMigrate(&models.Category{}, &models.Product{})

	if err != nil {
		fmt.Println("gagal migrate", err)
	}
	DB = database
}
