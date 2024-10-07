package route

import (
	"github.com/gin-gonic/gin"
	"shop/order_service/inner_layer/repository"
	"shop/order_service/outer_layer/adapter"
)

const (
	APP_GROUP       = "/v1"
	ORDER_GROUP     = "/order"
	ORDER_ADD_PATH  = "/add"
	ORDER_PAY_PATH  = "/pay"
	ALL_ORDERS_PATH = "/all"
)

func ApplicationRoutes(routes *gin.Engine) {
	group := routes.Group(APP_GROUP)

	bAdapter := adapter.BaseAdapter{
		Repository: repository.Repository{},
	}

	OrdersRoutes(group, bAdapter.OrdersAdapter())
}
