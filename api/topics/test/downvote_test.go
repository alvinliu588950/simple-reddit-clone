package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/alvinliu588950/simple-reddit-clone/api/topics"
)

type DownoteResponse struct {
	Message string
}

func TestDownvoteRoute(t *testing.T) {
	r, topicRepo := setup()

	votes := 1
	id := 0
	topicRepo.Add(topics.Topic{
		Id: id,
		Content: "test",
		Votes: 1,
	})
	// mock request body
	payload := map[string]interface{}{
		"id": id,
	}
	body, _ := json.Marshal(payload)
	
	// make request
	req, _ := http.NewRequest("PATCH", "/api/v1/topics/downvote", bytes.NewReader(body))
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Convert the JSON response
	var response AddResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// assertions on the correctness of the response.
	topic, _ := topicRepo.Find(id)
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "topic downvote success", response.Message)
	assert.Equal(t, votes - 1, topic.Votes)
}