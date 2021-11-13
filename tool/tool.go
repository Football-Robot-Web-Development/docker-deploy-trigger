package tool

import (
	"github.com/gin-gonic/gin"
)


func RenderErrJson(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"code":    -1,
		"message": err.Error(),
	})
}
