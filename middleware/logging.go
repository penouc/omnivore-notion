package middleware

import (
	"bytes"
	"context"
	"os"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var url = os.Getenv("KV_URL")
var opts, err = redis.ParseURL(url)
if err != nil {
	panic(err)
}
var rdb = redis.NewClient(opts)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Wrap the ResponseWriter
		writer := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// 打印请求信息
		reqBody, _ := c.GetRawData()
		fmt.Printf("[INFO] Request: %s %s %s\n", c.Request.Method, c.Request.RequestURI, reqBody)

		err := rdb.Set(ctx, "hahaha", reqBody, 0).Err()
		if err != nil {
			panic(err)
		}
		// Process request
		c.Next()

		// Save end time
		end := time.Now()
		latency := end.Sub(start)

		respBody := writer.body.String()

		fmt.Printf("[INFO] Response: %s %s %s (%v)\n", c.Request.Method, c.Request.RequestURI, respBody, latency)
	}
}
