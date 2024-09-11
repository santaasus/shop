package adapter

import (
	userRepository "shop/user_service/inner_layer/repository/user"
	authService "shop/user_service/inner_layer/service/auth"
	authController "shop/user_service/outer_layer/rest/controller/auth"
)

func AuthAdapter() *authController.Controller {
	repository := userRepository.Repository{}
	service := authService.Service{UserRepository: repository}
	controller := authController.Controller{AuthService: &service}
	return &controller
}
