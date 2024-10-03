package service

import (
	"errors"
	domain "shop/products_service/inner_layer/domain/products"
	"shop/products_service/inner_layer/repository"

	domainErrors "github.com/santaasus/errors-handler"
)

type Service struct {
	Repository repository.IRepository
}

func (s *Service) GetProducts() (*[]domain.Product, error) {
	products, err := s.Repository.GetProducts()
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("no products"),
			Type: domainErrors.InternalServerError,
		}
	}

	return products, nil
}
