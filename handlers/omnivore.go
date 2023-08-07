package handlers

import (
	"github.com/gin-gonic/gin"
)

func OmnivoreHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}
