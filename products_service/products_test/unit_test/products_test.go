package unittest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	domain "shop/products_service/inner_layer/domain/products"
	"shop/products_service/outer_layer/adapter"
	route "shop/products_service/outer_layer/route"
)

type mockRepository struct {
}

// Mock db logic
func (mockRepository) GetProducts() (*[]domain.Product, error) {
	var products *[]domain.Product = &[]domain.Product{{ID: 1, ProductName: ""}}

	return products, nil
}

func TestLogin(t *testing.T) {
	engine := gin.Default()
	group := engine.Group(route.APP_GROUP)

	bAdapter := &adapter.BaseAdapter{
		Repository: mockRepository{},
	}
	route.ProductsRoutes(group, bAdapter.ProductsAdapter())

	newRequest, err := http.NewRequest("GET", route.APP_GROUP+route.PRODUCTS_GROUP+"/", nil)

	if err != nil {
		t.Error(err)
		return
	}

	writer := httptest.NewRecorder()
	engine.ServeHTTP(writer, newRequest)

	assert.Equal(t, http.StatusOK, writer.Code)
}
