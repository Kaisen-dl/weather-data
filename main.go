package main

import (
	"fmt"

	"github.com/Kaisen-dl/weather-data/server"
	"github.com/gin-gonic/gin"
)

const (
	serverPort = ":9091"
)

func main() {
	fmt.Println("Я вова мне похуй!")

	router := gin.Default()
	ginHandlers := server.NewGinHandler()
	server.RegisterHandlers(router, ginHandlers)
	if err := router.Run(serverPort); err != nil {
		panic(err)
	}
}
