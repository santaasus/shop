package repository

import (
	"encoding/json"
	"errors"
	rDB "shop/db_service/redis"
	"shop/order_service/inner_layer/db"
	domain "shop/order_service/inner_layer/domain/order"
	"strconv"

	domainErrors "github.com/santaasus/errors-handler"
)

type IRepository interface {
	GetOrders(userId int, isFromCache bool) (*[]domain.Order, error)
	GetOrderById(id, userId int) (*domain.Order, error)
	AddOrder(productId, userId int) (*domain.Order, error)
	PayOrder(id int) error
}

type Repository struct {
}

func (Repository) GetOrders(userId int, isFromCache bool) (*[]domain.Order, error) {
	if isFromCache {
		value, _ := rDB.GetValueBy(rDB.ORDER + strconv.Itoa(userId))
		if value != "" {
			var orders []domain.Order
			err := json.Unmarshal([]byte(value), &orders)

			return &orders, err
		}
	}

	orders, err := db.GetOrders(userId)
	if err != nil {
		return nil, err
	}

	rDB.SaveBy(rDB.ORDER+strconv.Itoa(userId), orders)

	return orders, nil
}

func (Repository) GetOrderById(id, userId int) (*domain.Order, error) {
	order, err := db.GetOrderById(id, userId)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (Repository) AddOrder(productId int, userId int) (*domain.Order, error) {
	isExists, err := db.IsExistsOrder(productId, userId)
	if err != nil {
		return nil, err
	}

	if isExists {
		return nil, &domainErrors.AppError{
			Err:  errors.New("the order is exists"),
			Type: domainErrors.ValidationError,
		}
	}

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
