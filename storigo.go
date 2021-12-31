package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoha/storigo/config"
)

func main() {
	// read the dot env configs
	config.ReadConfigs()

	// init http server ( using gin )
	initHttpServer()
}

func initHttpServer() {
	gin.SetMode(config.Get(config.GIN_MODE))
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "app is running",
		})
	})

	server.Run(config.Get(config.APP_HOST_AND_PORT))
}
