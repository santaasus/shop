package service

import (
	"errors"
	domain "shop/order_service/inner_layer/domain/order"
	"shop/order_service/inner_layer/repository"

	jwtHandler "github.com/santaasus/JWTToken-handler"

	domainErrors "github.com/santaasus/errors-handler"
)

type Service struct {
	Repository repository.IRepository
}

func (s *Service) GetOrders(token string) (*[]domain.Order, error) {
	claims, err := jwtHandler.VerifyTokenAndGetClaims(token, jwtHandler.Access)
	if err != nil {
		return nil, err
	}

	userId := int(claims["id"].(float64))
	if userId == 0 {
		return nil, &domainErrors.AppError{
			Err:  errors.New("token meta info validate error"),
			Type: domainErrors.NotFound,
		}
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
	order, err := s.Repository.GetOrderById(id)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("no order"),
			Type: domainErrors.NotFound,
		}
	}

	return order, nil
}

func (s *Service) AddOrder(token string, productId int) (*domain.Order, error) {
	claims, err := jwtHandler.VerifyTokenAndGetClaims(token, jwtHandler.Access)
	if err != nil {
		return nil, err
	}

	userId := claims["id"].(int)
	if userId == 0 {
		return nil, &domainErrors.AppError{
			Err:  errors.New("token meta info validate error"),
			Type: domainErrors.NotFound,
		}
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

func (s *Service) PayOrder(id int) error {
	err := s.Repository.PayOrder(id)
	if err != nil {
		return &domainErrors.AppError{
			Err:  errors.New("something went wrong"),
			Type: domainErrors.ValidationError,
		}
	}

	return nil
}
