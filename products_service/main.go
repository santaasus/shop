package productsservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"shop/user_service/outer_layer/rest/route"

	logger "github.com/santaasus/logger-handler"
	db "shop/db_service"

	errorHandler "github.com/santaasus/errors-handler"

	limit "github.com/aviddiviner/gin-limit"
)

type ServerConfig struct {
	ServerPort int `json:"ServerPort"`
}

func main() {
	router := gin.Default()
	// https://github.com/aviddiviner/gin-limit/blob/master/README.md
	router.Use(limit.MaxAllowed(runtime.GOMAXPROCS(0) / 2))
	router.Use(cors.Default())
	router.Use(logger.GinBodyLogMiddleware)
	router.Use(errorHandler.ErrorHandler)

	_, err := db.Connect()
	if err != nil {
		_ = fmt.Errorf("fatal error in database file: %s", err)
		panic(err)
	}

	route.ApplicationRoutes(router)
	startServer(router)
}

func startServer(router http.Handler) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		_ = fmt.Errorf("error for open: %s", err.Error())
		panic(err)
	}

	var config ServerConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		_ = fmt.Errorf("error for unmarshal: %s", err.Error())
		panic(err)
	}

	server := &http.Server{
		Addr:           ":8081", //fmt.Sprintf(":%d", config.ServerPort),
		Handler:        router,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: 1 << 10,
	}

	if err := server.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description: %s", err.Error())
		panic(err)
	}

	fmt.Printf("server was started %v", server)
}
