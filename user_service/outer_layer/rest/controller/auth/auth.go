package auth

import (
	"net/http"

	domainErrors "github.com/santaasus/errors-handler"
	service "shop/user_service/inner_layer/service/auth"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	AuthService *service.Service
}

func (c *Controller) Login(ctx *gin.Context) {
	var login LoginRequest

	if err := ctx.BindJSON(&login); err != nil {
		appError := domainErrors.ThrowAppErrorWith(domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	authModel, err := c.AuthService.Login(MapToDomainUser(&login))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, authModel)
}

func (c *Controller) GetAccessTokenBy(ctx *gin.Context) {
	var accessToken AccessTokenRequest

	if err := ctx.BindJSON(&accessToken); err != nil {
		appError := domainErrors.ThrowAppErrorWith(domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	tokenModel, err := c.AuthService.AccessTokenByRefreshToken(accessToken.RefreshToken)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, tokenModel)
}
