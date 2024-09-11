package route

import (
	controller "shop/user_service/outer_layer/rest/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controller *controller.Controller) {
	group := router.Group("/user")
	{
		group.POST("/create", controller.CreateUser)
	}
}
