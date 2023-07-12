package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		id, _ := os.Hostname()
		host := c.Request.Host
		c.String(http.StatusOK, fmt.Sprintf("hello from go-gin , the host address is: %s, and the hostname is: %s", host, id))
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
