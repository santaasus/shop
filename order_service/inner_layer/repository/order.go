package repository

import (
	"shop/order_service/inner_layer/db"
	domain "shop/order_service/inner_layer/domain/order"
)

type IRepository interface {
	GetOrders(userId int) (*[]domain.Order, error)
	GetOrderById(id int) (*domain.Order, error)
	AddOrder(productId int, userId int) (*domain.Order, error)
	PayOrder(id int) error
}

type Repository struct {
}

func (Repository) GetOrders(userId int) (*[]domain.Order, error) {
	orders, err := db.GetOrders(userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (Repository) GetOrderById(id int) (*domain.Order, error) {
	order, err := db.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (Repository) AddOrder(productId int, userId int) (*domain.Order, error) {
	order, err := db.AddOrder(productId, userId)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (Repository) PayOrder(id int) error {
	err := db.PayOrder(id)
	if err != nil {
		return err
	}

	return nil
}
