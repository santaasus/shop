package adapter

import (
	"shop/products_service/inner_layer/service"
	"shop/products_service/outer_layer/controller"
)

func (b *BaseAdapter) ProductsAdapter() *controller.Controller {
	pService := service.Service{Repository: b.Repository}
	pController := controller.Controller{Service: &pService}
	return &pController
}
