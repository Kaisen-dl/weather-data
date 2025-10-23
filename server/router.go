package server

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, handler *GinHandler) {
	r.GET("/:city", handler.CityHandler)
}
