package route

import (
	"github.com/gin-gonic/gin"
	repository "shop/user_service/inner_layer/repository/user"
	adapter "shop/user_service/outer_layer/rest/adapter"
)

const (
	APP_GROUP        = "/v1"
	AUTH_GROUP       = "/auth"
	USER_GROUP       = "/user"
	EXIST_USER_GROUP = "/exist"
	LOGIN_PATH       = "/login"
	CREATE_USER_PATH = "/create"
	GET_USER_PATH    = "/"
	UPDATE_USER_PATH = "/update"
	DELETE_USER_PATH = "/delete"
)

func ApplicationRoutes(router *gin.Engine) {
	routerGroup := router.Group(APP_GROUP)
	baseAdapter := adapter.BaseAdapter{
		Repository: repository.Repository{},
	}

	AuthRoutes(routerGroup, baseAdapter.AuthAdapter())
	UserRoutes(routerGroup, baseAdapter.UserAdapter())
}
