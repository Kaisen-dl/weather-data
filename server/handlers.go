package server

import (
	"log"
	"net/http"

	"github.com/Kaisen-dl/weather-data/api"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	geocodingClient *api.Client
	omClient        *api.OMClient
}

func NewGinHandler(geocodingClient *api.Client, omClient *api.OMClient) *GinHandler {
	return &GinHandler{
		geocodingClient: geocodingClient,
		omClient:        omClient,
	}
}

func (g *GinHandler) CityHandler(c *gin.Context) {
	city := c.Param("city")
	log.Println(city)
	GeoResp, err := g.geocodingClient.GetCoordinates(city)
	if err != nil {
		log.Println("err:", err)
		return
	}

	omRes, err := g.omClient.GetTemperature(GeoResp.Latitude, GeoResp.Longitude)
	if err != nil {
		log.Println("err:", err)
		return
	}

	c.JSON(http.StatusOK, omRes)
}
