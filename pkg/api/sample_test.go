package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSample(t *testing.T) {
	router := SetupRouter()

	expectedBody := gin.H{
		"message": "sample JSON",
	}
	w := performRequest(router, "GET", "/api/v1/sample")
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedBody["message"], value)
}
