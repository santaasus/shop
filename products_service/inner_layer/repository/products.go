package repository

import (
	"shop/products_service/inner_layer/db"
	domain "shop/products_service/inner_layer/domain/products"
)

type IRepository interface {
	GetProducts() (*[]domain.Product, error)
}

type Repository struct {
}

func (Repository) GetProducts() (*[]domain.Product, error) {
	products, err := db.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
