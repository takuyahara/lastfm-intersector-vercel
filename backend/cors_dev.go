//go:build !prod

package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func config(router *gin.Engine) {
	fmt.Println("Running in non-prod environment...")
	useCors(router)
}

func useCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:3000",
			"http://localhost:3000",
		},
	}))
}
