package route

import (
	"shop/order_service/outer_layer/controller"

	"github.com/gin-gonic/gin"
)

func OrdersRoutes(routes *gin.RouterGroup, controller *controller.Controller) {
	group := routes.Group(ORDER_GROUP)
	{
		group.POST(ORDER_ADD_PATH, controller.AddOrder)
		group.PUT(ORDER_PAY_PATH, controller.PayOrder)
		group.GET(ALL_ORDERS_PATH, controller.GetOrders)
		group.GET("", controller.GetOrderById)
	}
}
