package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penouc/golang/handlers"
	"github.com/penouc/golang/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logging())

	r.GET("/ping", handlers.PingHandler)
	r.GET("/hello/:name", handlers.HelloHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
