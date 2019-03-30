package main

import (
	"github.com/gin-gonic/gin"
	"github.com/alvinliu588950/simple-reddit-clone/topics"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})		
	})

	r.Run()
}