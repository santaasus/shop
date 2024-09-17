package route

import (
	controller "shop/user_service/outer_layer/rest/controller/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routerGroup *gin.RouterGroup, controller *controller.Controller) {
	group := routerGroup.Group("/auth")
	{
		group.POST("/login", controller.Login)
	}
}
