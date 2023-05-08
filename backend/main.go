package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/takuyahara/lastfm-intersector-render/backend/api"
)

func main() {
	StartGin()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	config(router)
	api.Route(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
