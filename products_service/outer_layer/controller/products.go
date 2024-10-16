package controller

import (
	"net/http"
	"shop/products_service/inner_layer/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *service.Service
}

func (c *Controller) GetProducts(ctx *gin.Context) {
	cacheControl := ctx.GetHeader("Cache-Control")
	products, err := c.Service.GetProducts(cacheControl != "no-cache")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}
