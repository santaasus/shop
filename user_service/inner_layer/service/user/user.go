package user

import (
	userDomain "shop/user_service/inner_layer/domain/user"
	repository "shop/user_service/inner_layer/repository/user"
	security "shop/user_service/inner_layer/security"
)

type Service struct {
	Repository *repository.Repository
}

func (s *Service) CreateUser(user *userDomain.NewUser) (*userDomain.User, error) {
	domain := user.MapToDomainUser()

	hashPassword, err := security.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}

	domain.HashPassword = string(hashPassword)

	return s.Repository.CreateUser(domain)
}
