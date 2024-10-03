package route

import (
	"github.com/gin-gonic/gin"
	"shop/products_service/inner_layer/repository"
	"shop/products_service/outer_layer/adapter"
)

const (
	APP_GROUP      = "/v1"
	PRODUCTS_GROUP = "/products"
)

func ApplicationRoutes(routes *gin.Engine) {
	group := routes.Group(APP_GROUP)

	bAdapter := adapter.BaseAdapter{
		Repository: repository.Repository{},
	}

	ProductsRoutes(group, bAdapter.ProductsAdapter())
}
