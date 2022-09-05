package router

import (
	"net/http"

	"spotify-server/app/spotify"

	"github.com/gin-gonic/gin"
)

// spotifySwipeTokenDetails 交换spotify access token接口的请求body
type spotifySwipeTokenDetails struct {
	Code        string `form:"code"`
	RedirectURL string `form:"redirect_url"` // 客户端在query参数中设定
}

// SwipeSpotifyToken 使用auth code换取access_token, refresh_token
func SwipeSpotifyToken(c *gin.Context) {

	var body spotifySwipeTokenDetails

	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 调用Spotify API
	token := spotify.SwipeToken(body.Code, body.RedirectURL)
	c.JSON(http.StatusOK, token)
}

// spotifyRefreshTokenDetails 刷新spotify token接口的请求body
type spotifyRefreshTokenDetails struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

// RefreshSpotifyToken 刷新spotify token
func RefreshSpotifyToken(c *gin.Context) {
	var body spotifyRefreshTokenDetails

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 调用Spotify API
	token := spotify.RefreshToken(body.RefreshToken)
	c.JSON(http.StatusOK, token)
}
