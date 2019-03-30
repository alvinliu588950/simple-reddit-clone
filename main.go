package main

import (
	"github.com/gin-gonic/gin"
	"github.com/alvinliu588950/simple-reddit-clone/api/topics"
)

func main() {
	r := gin.Default()

	// in-memory storage
	topicRepo := topics.NewRepo()

	// Setup route group for the API
	v1 := r.Group("/api/v1")
	// register topics routes
	topics.Routes(v1, topicRepo)

	r.Run()
}