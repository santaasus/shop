package adapter

import (
	repository "shop/user_service/inner_layer/repository/user"
	service "shop/user_service/inner_layer/service/user"
	controller "shop/user_service/outer_layer/rest/controller/user"
)

func UserAdapter() *controller.Controller {
	var repository repository.Repository
	service := service.Service{Repository: &repository}
	controller := controller.Controller{Service: &service}

	return &controller
}
