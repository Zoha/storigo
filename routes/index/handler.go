package index

import (
	"github.com/gin-gonic/gin"
)

var IndexHandler gin.HandlerFunc = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "app is running",
	})
}
