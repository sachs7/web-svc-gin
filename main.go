package main

import (
	"fmt"
	"web-service-gin/client"
	"web-service-gin/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", server.GetAlbums)
	router.GET("/albums/:id", server.GetAlbumByID)
	router.POST("/albums", server.PostAlbums)
	router.GET("/readiness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "readiness - UP",
		})
	})
	router.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "liveness - LIVE",
		})
	})

	router.Run("0.0.0.0:8082")

	album, err := client.GetUserByID("0.0.0.0:8082", "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(album)
}
