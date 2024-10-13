package route

import (
	"github.com/gin-gonic/gin"
	"shop/order_service/outer_layer/controller"
	"shop/order_service/outer_layer/middleware"
)

func OrdersRoutes(routes *gin.RouterGroup, controller *controller.Controller) {
	group := routes.Group(ORDER_GROUP)
	group.Use(middleware.ValidateJWTMiddleware())
	{
		group.POST(ORDER_ADD_PATH, controller.AddOrder)
		group.PUT(ORDER_PAY_PATH, controller.PayOrder)
		group.GET(ALL_ORDERS_PATH, controller.GetOrders)
		group.GET("", controller.GetOrderById)
	}
}
