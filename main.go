package main

import (
	"image-retrieval/internal/resource/database"
	"image-retrieval/internal/resource/database/model"
	"image-retrieval/internal/resource/es"
	"image-retrieval/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	es.ConnectEs()
	// args := os.Args
	// if len(args) == 2 && args[1] == "migrage" {

	// }
	// fmt.Println("Migrating database...")
	model.InitModel()
	es.CreateIndex(&model.Image{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("/images", "./static/images")
	router.InitUserRouter(r)
	router.InitImageRouter(r)
	r.Run(":3000")
}
