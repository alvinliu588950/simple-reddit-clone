package topics

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {}

func (th *TopicHandler) List(tr *TopicRepo) gin.HandlerFunc {
	return func (c *gin.Context) {
		topics := tr.GetAll()

		c.JSON(http.StatusOK, gin.H {
			"topics": topics,
		})
	}
}
