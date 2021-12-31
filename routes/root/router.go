package root

import (
	"github.com/gin-gonic/gin"
	"github.com/zoha/storigo/routes/common"
)

func RootRouterInit(r *gin.RouterGroup) {
	r.GET("/", RootIndexHandler)
}

func RootIndexHandler(c *gin.Context) {
	rootResponse := common.MessageResponse{
		Message: "app is running",
	}
	c.JSON(200, rootResponse)
}
