package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoha/storigo/config"
	"github.com/zoha/storigo/routes/index"
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
	IndexGroup := server.Group("/")
	index.IndexInit(IndexGroup)

	server.Run(config.Get(config.APP_HOST_AND_PORT))
}
