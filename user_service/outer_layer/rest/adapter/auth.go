package adapter

import (
	authService "shop/user_service/inner_layer/service/auth"
	authController "shop/user_service/outer_layer/rest/controller/auth"
)

func (a *BaseAdapter) AuthAdapter() *authController.Controller {
	service := authService.Service{UserRepository: a.Repository}
	controller := authController.Controller{AuthService: &service}
	return &controller
}
