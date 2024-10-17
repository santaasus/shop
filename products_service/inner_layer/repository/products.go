package repository

import (
	"encoding/json"
	rDB "shop/db_service/redis"
	"shop/products_service/inner_layer/db"
	domain "shop/products_service/inner_layer/domain/products"
)

type IRepository interface {
	GetProducts(isFromCache bool) (*[]domain.Product, error)
}

type Repository struct {
}

func (Repository) GetProducts(isFromCache bool) (*[]domain.Product, error) {
	if isFromCache {
		value, _ := rDB.GetValueBy(rDB.PRODUCTS)
		if value != "" {
			var products []domain.Product
			err := json.Unmarshal([]byte(value), &products)

			return &products, err
		}
	}

	products, err := db.GetProducts()
	if err != nil {
		return nil, err
	}

	rDB.SaveBy(rDB.PRODUCTS, products)

	return products, nil
}
