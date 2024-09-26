package route

import (
	controller "shop/user_service/outer_layer/rest/controller/user"
	"shop/user_service/outer_layer/rest/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controller *controller.Controller) {
	group := router.Group(USER_GROUP)
	JWTGroup := group.Group(EXIST_USER_GROUP)
	JWTGroup.Use(middleware.AuthJWTMiddleware())
	{
		group.POST(CREATE_USER_PATH, controller.CreateUser)
		JWTGroup.GET(GET_USER_PATH+":id", controller.GetUser)
		JWTGroup.PUT(UPDATE_USER_PATH+"/:id", controller.UpdateUser)
		JWTGroup.DELETE(DELETE_USER_PATH+"/:id", controller.DeleteUser)
	}
}
