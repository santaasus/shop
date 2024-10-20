package adapter

import (
	"shop/notification_service/inner_layer/service"
	"shop/notification_service/outer_layer/controller"
)

func (b *BaseAdapter) NotificationAdapter() *controller.Controller {
	pService := service.Service{Repository: b.Repository}
	pController := controller.Controller{Service: &pService}
	return &pController
}
