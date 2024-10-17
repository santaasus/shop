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

func (s *Service) GetProducts(isFromCache bool) (*[]domain.Product, error) {
	products, err := s.Repository.GetProducts(isFromCache)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("no products"),
			Type: domainErrors.NotFound,
		}
	}

	return products, nil
}
