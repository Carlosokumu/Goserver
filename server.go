package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server.go/controllers"
	"server.go/models"
)

func main() {
	//Initialize the "Gin" router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	//Create a  database  connection  in the main Goroutine
	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.Run()
}
