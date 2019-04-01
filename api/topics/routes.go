/*
 *  Wirtten by the author AlvinLiu.
 */
package topics

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, tr *TopicRepo) {
	handler := TopicHandler{}
	r.GET("/topics", handler.List(tr))
	r.POST("/topics/add", handler.Add(tr))
	r.PATCH("/topics/upvote", handler.Upvote(tr))
	r.PATCH("/topics/downvote", handler.Downvote(tr))
}
