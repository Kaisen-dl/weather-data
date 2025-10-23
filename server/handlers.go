package server

import (
	"log"
	"net/http"

	"github.com/Kaisen-dl/weather-data/api"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	geocodingClient *api.Client
}

func NewGinHandler(geocodingClient *api.Client) *GinHandler {
	return &GinHandler{
		geocodingClient: geocodingClient,
	}
}

func (g *GinHandler) CityHandler(c *gin.Context) {
	city := c.Param("city")
	log.Println(city)
	resp, err := g.geocodingClient.GetCoordinates(city)
	if err != nil {
		log.Println("err:", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
