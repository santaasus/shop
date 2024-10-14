package unittests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	domain "shop/order_service/inner_layer/domain/order"
	"shop/order_service/outer_layer/adapter"
	"shop/order_service/outer_layer/route"
	"strings"
	"testing"

	structures "shop/order_service/outer_layer/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
}

type ginConfig struct {
	engine *gin.Engine
	group  *gin.RouterGroup
}

type requestConfig struct {
	method  string
	url     string
	reqBody any
	header  map[string]string
}

func (mockRepository) GetOrderById(id int) (*domain.Order, error) {
	order := &domain.Order{
		ID:        1,
		UserId:    1,
		ProductId: 1,
		IsPayed:   false,
	}

	return order, nil
}

func (mockRepository) GetOrders(userId int) (*[]domain.Order, error) {
	orders := &[]domain.Order{{
		ID:        1,
		UserId:    1,
		ProductId: 1,
		IsPayed:   false,
	}}

	return orders, nil
}

func (mockRepository) PayOrder(id int) error {
	return nil
}

func (mockRepository) AddOrder(productId, userId int) (*domain.Order, error) {
	order := &domain.Order{
		ID:        1,
		UserId:    1,
		ProductId: 1,
		IsPayed:   false,
	}

	return order, nil
}

// Change to the root path for correct reading files like os.Readfile
func changeWd() {
	path, _ := os.Getwd()
	rootPath := strings.Split(path, "/shop")[0]
	if len(rootPath) > 0 {
		os.Chdir(rootPath + "/shop")
	}
}

func getGinConfig() *ginConfig {
	engine := gin.Default()
	group := engine.Group(route.APP_GROUP)

	return &ginConfig{
		engine: engine,
		group:  group,
	}
}

func (c *requestConfig) doRequest(gin *ginConfig) (*httptest.ResponseRecorder, error) {
	jsonValue, err := json.Marshal(c.reqBody)
	if err != nil {
		return nil, err
	}

	newRequest, err := http.NewRequest(c.method, c.url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	if len(c.header) > 0 {
		for key, value := range c.header {
			newRequest.Header.Add(key, value)
		}
	}

	writer := httptest.NewRecorder()
	gin.engine.ServeHTTP(writer, newRequest)

	return writer, nil
}

func TestGetOrderById(t *testing.T) {
	changeWd()

	bAdapter := adapter.BaseAdapter{
		Repository: mockRepository{},
	}
	ginConfig := getGinConfig()
	route.OrdersRoutes(ginConfig.group, bAdapter.OrdersAdapter())

	reqConfig := requestConfig{
		method:  "GET",
		url:     route.APP_GROUP + route.ORDER_GROUP + fmt.Sprintf("?id=%d", 1),
		reqBody: nil,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyOTIxMzQyMX0.dy1jjqq3W_CLMGxwLcWm2OAevva6k7X0aN8zr2DKUE0"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetOrders(t *testing.T) {
	changeWd()

	bAdapter := adapter.BaseAdapter{
		Repository: mockRepository{},
	}
	ginConfig := getGinConfig()
	route.OrdersRoutes(ginConfig.group, bAdapter.OrdersAdapter())

	reqConfig := requestConfig{
		method:  "GET",
		url:     route.APP_GROUP + route.ORDER_GROUP + route.ALL_ORDERS_PATH,
		reqBody: nil,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyOTIxMzQyMX0.dy1jjqq3W_CLMGxwLcWm2OAevva6k7X0aN8zr2DKUE0"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestPayOrder(t *testing.T) {
	changeWd()

	bAdapter := adapter.BaseAdapter{
		Repository: mockRepository{},
	}
	ginConfig := getGinConfig()
	route.OrdersRoutes(ginConfig.group, bAdapter.OrdersAdapter())

	reqBody := &structures.PayOrderRequest{
		OrderId: 1,
	}

	reqConfig := requestConfig{
		method:  "PUT",
		url:     route.APP_GROUP + route.ORDER_GROUP + route.ORDER_PAY_PATH,
		reqBody: reqBody,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyOTIxMzQyMX0.dy1jjqq3W_CLMGxwLcWm2OAevva6k7X0aN8zr2DKUE0"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, http.StatusOK, writer.Code)

	response := `{"success":true}`

	assert.Equal(t, response, writer.Body.String())
}

func TestAddOrder(t *testing.T) {
	changeWd()

	bAdapter := adapter.BaseAdapter{
		Repository: mockRepository{},
	}
	ginConfig := getGinConfig()
	route.OrdersRoutes(ginConfig.group, bAdapter.OrdersAdapter())

	reqBody := &structures.AddOrderRequest{
		ProductId: 1,
	}

	reqConfig := requestConfig{
		method:  "POST",
		url:     route.APP_GROUP + route.ORDER_GROUP + route.ORDER_ADD_PATH,
		reqBody: reqBody,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyOTIxMzQyMX0.dy1jjqq3W_CLMGxwLcWm2OAevva6k7X0aN8zr2DKUE0"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}
