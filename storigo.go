package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoha/storigo/config"
	rootRouter "github.com/zoha/storigo/routes/root"
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
	index := server.Group("/")
	rootRouter.RootRouterInit(index)

	server.Run(config.Get(config.APP_HOST_AND_PORT))
}
