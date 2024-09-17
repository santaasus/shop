package user

import (
	"errors"
	"net/http"
	domainErrors "shop/user_service/inner_layer/domain/errors"
	userService "shop/user_service/inner_layer/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *userService.Service
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var request NewUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	userModel, err := c.Service.CreateUser(MapToDomainUser(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, userModel)
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the id type"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	var request UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the request body"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	err = c.Service.UpdateUser(MapToDomainUpdateUser(request), userId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the id type"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	err = c.Service.DeleteUser(userId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (c *Controller) GetUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the id type"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	user, err := c.Service.GetUser(userId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &user)
}
