package main

import (
	"golang/bmysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginApp := gin.Default()
	bmysql.InitializeConnection()
	ginApp.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	ginApp.GET("/albums/all", getAllAlbums)
	// listen on port:8080
	ginApp.Run()
}

func getAllAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, bmysql.GetAllAlbums())
}
