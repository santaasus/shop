package controller

import (
	"errors"
	"net/http"
	"shop/order_service/inner_layer/service"
	"strconv"

	"github.com/gin-gonic/gin"
	domainErrors "github.com/santaasus/errors-handler"
)

type Controller struct {
	Service *service.Service
}

func (c *Controller) GetOrders(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	orders, err := c.Service.GetOrders(token)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *Controller) GetOrderById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the id type"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	token := ctx.GetHeader("Authorization")
	order, err := c.Service.GetOrderById(token, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *Controller) AddOrder(ctx *gin.Context) {
	var request AddOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the request body"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	token := ctx.GetHeader("Authorization")
	order, err := c.Service.AddOrder(token, request.ProductId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *Controller) PayOrder(ctx *gin.Context) {
	var request PayOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err = &domainErrors.AppError{
			Err:  errors.New("wrong the request body"),
			Type: domainErrors.ValidationError,
		}
		_ = ctx.Error(err)
		return
	}

	err := c.Service.PayOrder(request.OrderId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
