package route

import (
	"github.com/gin-gonic/gin"
	"shop/products_service/outer_layer/controller"
)

func ProductsRoutes(routes *gin.RouterGroup, controller *controller.Controller) {
	group := routes.Group(PRODUCTS_GROUP)
	{
		group.GET("/", controller.GetProducts)
	}
}
