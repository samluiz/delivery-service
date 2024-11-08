package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONResponse(t *testing.T) {
	data := map[string]string{"message": "success"}

	recorder := httptest.NewRecorder()

	NewJSONResponse(recorder, http.StatusOK, data)

	assert.Equal(t, http.StatusOK, recorder.Code)

	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	var responseBody map[string]string
	err := json.NewDecoder(recorder.Body).Decode(&responseBody)
	assert.Nil(t, err)
	assert.Equal(t, data, responseBody)
}
