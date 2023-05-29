package main

import (
	"fmt"
	"net/http"

	"github.com/xingxingso/gin"
	//"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping",
		func(c *gin.Context) {
			fmt.Println("is coming")
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		},
		func(c *gin.Context) {
			fmt.Println("is done")
		},
	)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
