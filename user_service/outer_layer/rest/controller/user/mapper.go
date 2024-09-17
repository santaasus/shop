package user

import (
	userDomain "shop/user_service/inner_layer/domain/user"
)

func MapToDomainUser(request *NewUserRequest) *userDomain.NewUser {
	return &userDomain.NewUser{
		Email:     request.Email,
		Password:  request.Password,
		UserName:  request.UserName,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
}

func MapToDomainUpdateUser(request UpdateUserRequest) userDomain.UpdateUser {
	return userDomain.UpdateUser{
		Email:     request.Email,
		UserName:  request.UserName,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
}
