package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"shop/user_service/outer_layer/rest/route"

	logger "github.com/santaasus/logger-handler"
	db "shop/db_service"

	errorHandler "github.com/santaasus/errors-handler"

	limit "github.com/aviddiviner/gin-limit"
	root "shop"
)

type Config struct {
	ServerPort ServerPort `json:"ServerPort"`
}

type ServerPort struct {
	Port int `json:"user"`
}

func main() {
	router := gin.Default()
	// https://github.com/aviddiviner/gin-limit/blob/master/README.md
	router.Use(limit.MaxAllowed(runtime.GOMAXPROCS(0) / 2))
	router.Use(cors.Default())

	_, err := db.Connect()
	if err != nil {
		_ = fmt.Errorf("fatal error in database file: %s", err)
		panic(err)
	}

	router.Use(logger.GinBodyLogMiddleware)
	router.Use(errorHandler.ErrorHandler)
	route.ApplicationRoutes(router)
	startServer(router)
}

func startServer(router http.Handler) {
	data, err := root.FileByName("config.json")
	if err != nil {
		_ = fmt.Errorf("error for open: %s", err.Error())
		panic(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		_ = fmt.Errorf("error for unmarshal: %s", err.Error())
		panic(err)
	}

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerPort.Port),
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
