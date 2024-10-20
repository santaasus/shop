package unittests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	route "shop/notification_service/outer_layer/router"
	"strings"
	"testing"

	smtpCore "shop/notification_service/inner_layer/smtp"
	structures "shop/notification_service/outer_layer/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Change to the root path for correct reading files like os.Readfile
func changeWd() {
	path, _ := os.Getwd()
	rootPath := strings.Split(path, "/shop")[0]
	if len(rootPath) > 0 {
		os.Chdir(rootPath + "/shop")
	}
}

func TestNotifico(t *testing.T) {
	changeWd()

	go smtpCore.StartListenServ()

	engine := gin.Default()

	route.ApplicationRoutes(engine)

	reqBody := &structures.NotificationRequest{
		OrderId:   1,
		UserEmail: "test@gmail.com",
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		t.Error(err)
		return
	}

	newRequest, err := http.NewRequest("POST", route.APP_GROUP+route.EMAIL_GROUP+route.EMAIL_SEND_PATH, bytes.NewBuffer(data))

	if err != nil {
		t.Error(err)
		return
	}

	writer := httptest.NewRecorder()
	engine.ServeHTTP(writer, newRequest)

	assert.Equal(t, http.StatusOK, writer.Code)
}
