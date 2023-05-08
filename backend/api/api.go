package api

import (
	"github.com/gin-gonic/gin"
	"github.com/takuyahara/lastfm-intersector-render/backend/api/similarartists"
)

func Route(router *gin.Engine) {
	router.GET("/api/artist/:artist", similarartists.ArtistGET)
}
