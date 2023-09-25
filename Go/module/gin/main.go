package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/ping", pingHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	_ = r.Run(":1112")
}

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
