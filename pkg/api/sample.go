package api

import (
	"github.com/gin-gonic/gin"
)

// Sample REST API endpoint
func Sample(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "sample JSON",
	})
}