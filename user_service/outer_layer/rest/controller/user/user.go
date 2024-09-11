package user

import (
	"net/http"
	userService "shop/user_service/inner_layer/service/user"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *userService.Service
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var newUser NewUserRequest

	if err := ctx.BindJSON(&newUser); err != nil {
		_ = ctx.Error(err)
		return
	}

	userModel, err := c.Service.CreateUser(MapToDomainUser(&newUser))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, userModel)
}
