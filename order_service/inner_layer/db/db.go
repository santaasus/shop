package db

import (
	"fmt"
	dbCore "shop/db_service"
	domain "shop/order_service/inner_layer/domain/order"
)

func GetOrders(userId int) (*[]domain.Order, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT * FROM orders WHERE user_id=$1;"
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	orders := []domain.Order{}

	if rows.Next() {
		order := domain.Order{}
		rows.Scan(&order.ID, &order.UserId, &order.ProductId, &order.IsPayed)
		orders = append(orders, order)
	}

	return &orders, nil
}

func GetOrderById(id int) (*domain.Order, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT * FROM orders WHERE id=$1"
	row := db.QueryRow(query, id)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	var order = &domain.Order{}
	if err := row.Scan(&order.ID, &order.UserId, &order.ProductId, &order.IsPayed); err != nil {
		return nil, err
	}

	return order, nil
}

func AddOrder(productId, userId int) (*domain.Order, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var orderId = 0
	query := "INSERT INTO orders(user_id, product_id, is_payed) VALUES($1, $2, $3) RETURNING id;"
	err = db.QueryRow(query, userId, productId, false).Scan(&orderId)
	if err != nil {
		return nil, err
	}

	order := &domain.Order{
		ID:        orderId,
		UserId:    userId,
		ProductId: productId,
		IsPayed:   false,
	}

	return order, nil
}

func IsExistsOrder(productId, userId int) (bool, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return false, err
	}

	defer db.Close()

	var isExists bool
	query := "SELECT EXISTS(SELECT 1 FROM orders WHERE user_id=$1 AND product_id=$2);"
	_ = db.QueryRow(query, userId, productId).Scan(&isExists)

	return isExists, nil
}

func PayOrder(id int) error {
	db, err := dbCore.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	query := "UPDATE orders SET is_payed = true " + fmt.Sprintf(" WHERE id=%d;", id)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
