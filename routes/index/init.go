package index

import (
	"github.com/gin-gonic/gin"
)

func IndexInit(r *gin.RouterGroup) {
	r.GET("/", IndexHandler)
}
