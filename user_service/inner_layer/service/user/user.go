package user

import (
	userDomain "shop/user_service/inner_layer/domain/user"
	mapper "shop/user_service/inner_layer/repository/user"
	repository "shop/user_service/inner_layer/repository/user"
	security "shop/user_service/inner_layer/security"
)

type Service struct {
	Repository repository.IRepository
}

func (s *Service) CreateUser(user *userDomain.NewUser) (*userDomain.User, error) {
	domain := mapper.MapToDomainUser(user)

	hashPassword, err := security.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}

	domain.HashPassword = string(hashPassword)

	return s.Repository.CreateUser(domain)
}

func (s *Service) UpdateUser(updateUser userDomain.UpdateUser, userId int) error {
	err := s.Repository.UpdateUser(updateUser, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(userId int) error {
	_, err := s.GetUser(userId, false)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteUserByID(userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUser(userId int, isFromCache bool) (*userDomain.User, error) {
	user, err := s.Repository.GetUserByID(userId, isFromCache)
	if err != nil {
		return nil, err
	}

	return user, nil
}
