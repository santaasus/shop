package db

import (
	dbCore "shop/db_service"
	domain "shop/products_service/inner_layer/domain/products"
)

func GetProducts() (*[]domain.Product, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT * FROM products;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		rows.Scan(&product.ID, &product.ProductName)
		products = append(products, product)
	}

	return &products, nil
}
