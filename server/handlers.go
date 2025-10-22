package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
}

func NewGinHandler() *GinHandler {
	return &GinHandler{}
}

func (g *GinHandler) CityHandler(c *gin.Context) {
	city := c.Param("city")
	log.Println(city)
	c.JSON(http.StatusOK, gin.H{"message": "all good"})
}
