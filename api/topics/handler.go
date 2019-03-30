package topics

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {}

type AddRequest struct {
	Content string `form:"content" json:"content" xml:"content"  binding:"exists"`
}

type UpvoteRequest struct {
	Id int `form:"id" json:"id" xml:"id"  binding:"exists"`
}

type DownvoteRequest struct {
	Id int `form:"id" json:"id" xml:"id"  binding:"exists"`
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
		var body AddRequest

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

func (th *TopicHandler) Upvote(tr *TopicRepo) gin.HandlerFunc {
	return func (c *gin.Context) {
		var body UpvoteRequest

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tr.Upvote(body.Id)

		c.JSON(http.StatusOK, gin.H {
			"message": "topic upvote success",
		})
	}
}

func (th *TopicHandler) Downvote(tr *TopicRepo) gin.HandlerFunc {
	return func (c *gin.Context) {
		var body DownvoteRequest
		// validation
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		tr.Downvote(body.Id)
	
		c.JSON(http.StatusOK, gin.H {
			"message": "topic downvote success",
		})
	}


}
