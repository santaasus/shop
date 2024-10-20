package route

import (
	"github.com/gin-gonic/gin"
	"shop/notification_service/outer_layer/controller"
)

func NotificationRoutes(routes *gin.RouterGroup, controller *controller.Controller) {
	group := routes.Group(EMAIL_GROUP)
	{
		group.POST(EMAIL_SEND_PATH, controller.SendEmail)
	}
}
