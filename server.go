package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initialize the "Gin" router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
