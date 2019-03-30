package test

import (
	"github.com/gin-gonic/gin"
    "github.com/alvinliu588950/simple-reddit-clone/api/topics"
)

func setup() (*gin.Engine, *topics.TopicRepo) {
	r := gin.Default()
	// in-memory storage
	topicsRepo := topics.NewRepo()
	// Setup route group for the API
	v1 := r.Group("/api/v1")
	topics.Routes(v1, topicsRepo)

	// return router and storage
	return r, topicsRepo
}

