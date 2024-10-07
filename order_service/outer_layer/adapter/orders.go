package adapter

import (
	"shop/order_service/inner_layer/service"
	"shop/order_service/outer_layer/controller"
)

func (b *BaseAdapter) OrdersAdapter() *controller.Controller {
	pService := service.Service{Repository: b.Repository}
	pController := controller.Controller{Service: &pService}
	return &pController
}
