package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/notification_service/inner_layer/service"
)

type Controller struct {
	Service *service.Service
}

func (c *Controller) SendEmail(ctx *gin.Context) {
	var request NotificationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err)
		return
	}

	err := c.Service.SendMail(request.MapToDomain())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
