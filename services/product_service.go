package services

import (
	"golang_crud/models"
	"golang_crud/repositories"
	"golang_crud/utils"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (r *ProductService) FindAll(key []byte) ([]models.Product, error) {
	products, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range products {
		products[i].SecretNote = utils.Decrypt(key, products[i].SecretNote)
	}

	return products, nil
}
