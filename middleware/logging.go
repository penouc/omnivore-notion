package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		log.Printf("%s - %s %s %d %s",
			startTime.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL,
			c.Writer.Status(),
			time.Since(startTime),
		)
	}
}
