package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		remoteAddr := c.ClientIP()
		c.String(http.StatusOK, "it works on my machine @%s", remoteAddr)

	})

	router.GET("/time", func(c *gin.Context) {
		time := time.Now().Format(time.RFC3339)
		c.String(http.StatusOK, "Server time is: %s", time)
	})

	router.GET("/sayhello/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":8080")
}
