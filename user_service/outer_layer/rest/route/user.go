package route

import (
	controller "shop/user_service/outer_layer/rest/controller/user"
	"shop/user_service/outer_layer/rest/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controller *controller.Controller) {
	group := router.Group("/user")
	JWTGroup := group.Group("/exist")
	JWTGroup.Use(middleware.AuthJWTMiddleware())
	{
		group.POST("/create", controller.CreateUser)
		JWTGroup.GET("/:id", controller.GetUser)
		JWTGroup.PUT("/update/:id", controller.UpdateUser)
		JWTGroup.DELETE("/delete/:id", controller.DeleteUser)
	}
}
