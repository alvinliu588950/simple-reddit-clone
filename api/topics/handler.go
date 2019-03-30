package topics

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {}

type AddTopicRequest struct {
	Content string `form:"content" json:"content" xml:"content"  binding:"exists"`
}


func (th *TopicHandler) List(tr *TopicRepo) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"topics": tr.Topics,
		})
	}
}

func (th *TopicHandler) Add(tr *TopicRepo) gin.HandlerFunc {
	return func (c *gin.Context) {
		var body AddTopicRequest

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tr.Add(Topic{ 
			Id: len(tr.Topics),
			Content: body.Content,
			Votes: 0,
		})

		c.JSON(http.StatusOK, gin.H {
			"message": "topic add success",
		})
	}
}
