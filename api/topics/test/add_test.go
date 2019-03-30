package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

type AddResponse struct {
	Message string
}

func TestAddRoute(t *testing.T) {
	r, topicRepo := setup()

	// mock request body
	payload := map[string]interface{}{
		"content": "test add",
	}
	body, _ := json.Marshal(payload)
	
	// make request
	req, _ := http.NewRequest("POST", "/api/v1/topics/add", bytes.NewReader(body))
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Convert the JSON response
	var response AddResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "topic add success", response.Message)
	assert.Equal(t, 1, len(topicRepo.Topics))
	assert.Equal(t, topicRepo.Topics[0].Content, "test add")
}