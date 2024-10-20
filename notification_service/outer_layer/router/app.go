package route

import (
	"shop/notification_service/inner_layer/repository"
	"shop/notification_service/outer_layer/adapter"

	"github.com/gin-gonic/gin"
)

const (
	APP_GROUP       = "/v1"
	EMAIL_GROUP     = "/email"
	EMAIL_SEND_PATH = "/send"
)

func ApplicationRoutes(routes *gin.Engine) {
	group := routes.Group(APP_GROUP)

	bAdapter := adapter.BaseAdapter{
		Repository: repository.Repository{},
	}

	NotificationRoutes(group, bAdapter.NotificationAdapter())
}
