package topics

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, tr *TopicRepo) {
	handler := TopicHandler{}
	r.GET("/topics", handler.List(tr))
}
