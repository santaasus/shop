package route

import (
	"github.com/gin-gonic/gin"
	adapter "shop/user_service/outer_layer/rest/adapter"
)

func ApplicationRoutes(router *gin.Engine) {
	routerGroup := router.Group("/v1")
	AuthRoutes(routerGroup, adapter.AuthAdapter())
	UserRoutes(routerGroup, adapter.UserAdapter())
}
