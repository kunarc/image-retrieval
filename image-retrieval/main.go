package main

import (
	"image-retrieval/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.InitUserRouter(r)
	r.Run(":8080")
}
