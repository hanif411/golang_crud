package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name       string `json:"name"`
	Price      int    `json:"price"`
	SecretNote string `json:"secret_note"`
	CategoryID uint   `json:"category_id"`
}
