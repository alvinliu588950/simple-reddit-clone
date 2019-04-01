/*
 *  Wirtten by the author AlvinLiu.
 */
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/alvinliu588950/simple-reddit-clone/api/topics"
)

type ListResponse struct {
	Topics []topics.Topic
}

func TestListRoute(t *testing.T) {
	r, topicRepo := setup()
	// make request
	req, _ := http.NewRequest("GET", "/api/v1/topics", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Convert the JSON response
	var response ListResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	
	// assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, topicRepo.Topics, response.Topics)
}