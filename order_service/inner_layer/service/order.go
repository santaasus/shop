package service

import (
	"errors"
	"shop/order_service/inner_layer/repository"
	"shop/order_service/inner_layer/security"

	domain "shop/order_service/inner_layer/domain/order"

	domainErrors "github.com/santaasus/errors-handler"
)

type Service struct {
	Repository repository.IRepository
}

func (s *Service) GetOrders(token string, userId int) (*[]domain.Order, error) {
	err := security.ParseToken(token)
	if err != nil {
		return nil, err
	}

	orders, err := s.Repository.GetOrders(userId)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("no orders"),
			Type: domainErrors.NotFound,
		}
	}

	return orders, nil
}

func (s *Service) GetOrderById(token string, id int) (*domain.Order, error) {
	err := security.ParseToken(token)
	if err != nil {
		return nil, err
	}

	order, err := s.Repository.GetOrderById(id)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("no order"),
			Type: domainErrors.NotFound,
		}
	}

	return order, nil
}

func (s *Service) AddOrder(token string, productId, userId int) (*domain.Order, error) {
	err := security.ParseToken(token)
	if err != nil {
		return nil, err
	}

	order, err := s.Repository.AddOrder(productId, userId)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("something went wrong"),
			Type: domainErrors.ValidationError,
		}
	}

	return order, nil
}

func (s *Service) PayOrder(token string, id int) error {
	err := security.ParseToken(token)
	if err != nil {
		return err
	}

	err = s.Repository.PayOrder(id)
	if err != nil {
		return &domainErrors.AppError{
			Err:  errors.New("something went wrong"),
			Type: domainErrors.ValidationError,
		}
	}

	return nil
}
