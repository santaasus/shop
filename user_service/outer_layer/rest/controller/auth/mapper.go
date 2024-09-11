package auth

import (
	userDomain "shop/user_service/inner_layer/domain/user"
)

func MapToDomainUser(request *LoginRequest) *userDomain.LoginUser {
	return &userDomain.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}
}
