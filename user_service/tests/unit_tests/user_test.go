package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"shop/user_service/outer_layer/rest/controller/auth"
	"strconv"
	"strings"
	"testing"
	"time"

	"shop/user_service/outer_layer/rest/route"

	domain "shop/user_service/inner_layer/domain/user"
	structures "shop/user_service/inner_layer/service/auth"
	adapter "shop/user_service/outer_layer/rest/adapter"
	userController "shop/user_service/outer_layer/rest/controller/user"

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

var responseUser = &domain.User{
	ID:           6,
	UserName:     "Test",
	Email:        "Test@gmail.com",
	FirstName:    "Test",
	LastName:     "Test",
	HashPassword: "$2a$10$OD6gRRUd0O8cTxdGQjPzqOuk1cmMkX/FON.1jkfVpz0I.AuQXqvMa",
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
}

// Mock db logic
func (mockRepository) GetUserByID(id int) (*domain.User, error) {
	return responseUser, nil
}

func (mockRepository) GetUserByParams(params map[string]any) (*domain.User, error) {
	return responseUser, nil
}

func (mockRepository) CreateUser(newUser *domain.User) (*domain.User, error) {
	return responseUser, nil
}

func (mockRepository) UpdateUser(updateUser domain.UpdateUser, userId int) error {
	return nil
}

func (mockRepository) DeleteUserByID(userId int) error {
	return nil
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

func getAdapter() *adapter.BaseAdapter {
	return &adapter.BaseAdapter{
		Repository: mockRepository{},
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

func TestLogin(t *testing.T) {
	changeWd()

	request := &auth.LoginRequest{
		Email:    "Vladimirrrr@gmail.com",
		Password: "123456",
	}

	ginConfig := getGinConfig()
	route.AuthRoutes(ginConfig.group, getAdapter().AuthAdapter())

	reqConfig := requestConfig{
		method:  "POST",
		url:     route.APP_GROUP + route.AUTH_GROUP + route.LOGIN_PATH,
		reqBody: request,
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	// Check response status
	assert.Equal(t, http.StatusOK, writer.Code)

	var expected structures.AuthenticatedUser
	err = json.Unmarshal(writer.Body.Bytes(), &expected)
	if err != nil {
		t.Errorf("body len: %v, error for unmarshal: %v", len(writer.Body.Bytes()), err)
		return
	}

	assert.NotZero(t, expected.Data.ID)
}

func TestGetUser(t *testing.T) {
	changeWd()

	ginConfig := getGinConfig()
	route.UserRoutes(ginConfig.group, getAdapter().UserAdapter())

	reqConfig := requestConfig{
		method:  "GET",
		url:     route.APP_GROUP + route.USER_GROUP + route.EXIST_USER_GROUP + "/" + strconv.Itoa(responseUser.ID),
		reqBody: nil,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyNjUyNjE0MH0.CTRRvYEKCB2vj5AQ53js6SOGl7OGMvF65iMDXdodzBI"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	// Check response status
	assert.Equal(t, http.StatusOK, writer.Code)

	var expected domain.User
	err = json.Unmarshal(writer.Body.Bytes(), &expected)
	if err != nil {
		t.Errorf("body len: %v, error for unmarshal: %v", len(writer.Body.Bytes()), err)
		return
	}

	assert.NotZero(t, expected.ID)
}

func TestDeleteUser(t *testing.T) {
	changeWd()

	ginConfig := getGinConfig()
	route.UserRoutes(ginConfig.group, getAdapter().UserAdapter())

	reqConfig := requestConfig{
		method:  "DELETE",
		url:     route.APP_GROUP + route.USER_GROUP + route.EXIST_USER_GROUP + route.DELETE_USER_PATH + "/" + strconv.Itoa(responseUser.ID),
		reqBody: nil,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyNjUyNjE0MH0.CTRRvYEKCB2vj5AQ53js6SOGl7OGMvF65iMDXdodzBI"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	// Check response status
	assert.Equal(t, http.StatusOK, writer.Code)

	expected := `{"success":true}`
	assert.JSONEq(t, expected, writer.Body.String())
}

func TestUpdateUser(t *testing.T) {
	changeWd()

	ginConfig := getGinConfig()
	route.UserRoutes(ginConfig.group, getAdapter().UserAdapter())

	reqConfig := requestConfig{
		method:  "PUT",
		url:     route.APP_GROUP + route.USER_GROUP + route.EXIST_USER_GROUP + route.UPDATE_USER_PATH + "/" + strconv.Itoa(responseUser.ID),
		reqBody: nil,
		header:  map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTcyNjUyNjE0MH0.CTRRvYEKCB2vj5AQ53js6SOGl7OGMvF65iMDXdodzBI"},
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, http.StatusOK, writer.Code)

	expected := `{"success":true}`
	assert.JSONEq(t, expected, writer.Body.String())
}

func TestCreateUser(t *testing.T) {
	changeWd()

	ginConfig := getGinConfig()
	route.UserRoutes(ginConfig.group, getAdapter().UserAdapter())

	request := &userController.NewUserRequest{
		Email:     "asdsd@asdsad.ru",
		Password:  "123456",
		UserName:  "UserName",
		FirstName: "FirstName",
		LastName:  "LastName",
	}

	reqConfig := requestConfig{
		method:  "POST",
		url:     route.APP_GROUP + route.USER_GROUP + route.CREATE_USER_PATH,
		reqBody: request,
	}

	writer, err := reqConfig.doRequest(ginConfig)
	if err != nil {
		t.Error(err)
		return
	}

	// Check response status
	assert.Equal(t, http.StatusOK, writer.Code)

	var expected domain.User
	err = json.Unmarshal(writer.Body.Bytes(), &expected)
	if err != nil {
		t.Errorf("body len: %v, error for unmarshal: %v", len(writer.Body.Bytes()), err)
		return
	}

	assert.NotZero(t, expected.ID)
}
