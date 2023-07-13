package main

import (
//	"fmt"
	"context"
	"log"
	"net/http"
	"github.com/redis/go-redis/v9"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr:       "127.0.0.1:6379",
		Password:   "",
		DB:         0,
	})
	ctx := context.Background()

	/*
	   fmt.Println(rdb)

	   val, err := rdb.Do(ctx, "set", "hi", "sex").Text()
	   fmt.Println(val, err)
	*/

	router.GET("/set/:key/:value", func(c *gin.Context) {
		//remoteAddr := c.ClientIP()
		key := c.Param("key")
		value := c.Param("value")

		val, err := rdb.Do(ctx, "set", key, value).Text()

		if err != nil {
			log.Println("Error:", err)
			return
		}

		data := gin.H{
			"output": val,
		}

		c.JSON(http.StatusOK, data)
	})

	router.Run(":8080")
}

