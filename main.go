package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		id, _ := os.Hostname()
		host := c.Request.Host
		c.String(http.StatusOK, fmt.Sprintf("hello from node , the host address is: %s, and the hostname is: %s", host, id))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
