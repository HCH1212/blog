package api

import "github.com/gin-gonic/gin"

func Info(c *gin.Context) {
	c.JSON(200, gin.H{"mes": "1111"})
}
