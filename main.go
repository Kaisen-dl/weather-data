package main

import (
	"fmt"
	"sync"

	"github.com/Kaisen-dl/weather-data/cron_folder"
	"github.com/Kaisen-dl/weather-data/server"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
)

const (
	serverPort = ":9091"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("Я вова мне похуй!")

	//----------------------- SERVER ---------------------------
	router := gin.Default()
	ginHandlers := server.NewGinHandler()
	server.RegisterHandlers(router, ginHandlers)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Запуск сервера на порту" + serverPort)
		if err := router.Run(serverPort); err != nil {
			panic(err)
		}
	}()

	//----------------------- CRON -----------------------------
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	jobs, err := cron_folder.DoCron(s)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Запуск работы:", jobs[0].ID())
		s.Start()
	}()

	wg.Wait()
}
