package main

import (
  // "fmt"
  // "context"
   "net/http"
  // "github.com/redis/go-redis/v9"
  "github.com/gin-gonic/gin"
)

func main() {
   router := gin.Default()
/*
   rdb := redis.NewClient(&redis.Options{
        Addr: "127.0.0.1:6379",
	Password: "",
	DB: 0,
       })
   ctx := context.Background()


/*
   fmt.Println(rdb)

   val, err := rdb.Do(ctx, "set", "hi", "sex").Text()
   fmt.Println(val, err)
*/

     router.GET("/set", func(c *gin.Context) {
     //remoteAddr := c.ClientIP()
     value := c.Query("value")
     data := gin.H{
        "val": value,
	}

     c.JSON(http.StatusOK, data)
})

router.Run(":8080")
}
