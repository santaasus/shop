package adapter

import (
	service "shop/user_service/inner_layer/service/user"
	controller "shop/user_service/outer_layer/rest/controller/user"
)

func (a *BaseAdapter) UserAdapter() *controller.Controller {
	service := service.Service{Repository: a.Repository}
	controller := controller.Controller{Service: &service}
	return &controller
}
