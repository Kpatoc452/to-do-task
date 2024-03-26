package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello World",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", helloWorld)

	router.Run(":8080")
}
