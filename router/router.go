package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("token", SwipeSpotifyToken)
	r.POST("refresh_token", RefreshSpotifyToken)

	return r
}
