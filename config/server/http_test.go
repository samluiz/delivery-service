package server

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testes do servidor HTTP para garantir que ele está configurado corretamente

func TestNewServer(t *testing.T) {
	db := &sql.DB{}
	server := NewServer(db)

	assert.NotNil(t, server)
	assert.Equal(t, db, server.db)
	assert.NotNil(t, server.Router)
}

func TestServerServeHTTP(t *testing.T) {
	db := &sql.DB{}
	server := NewServer(db)

	server.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	req := httptest.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()

	server.ServeHTTP(recorder, req)

	res := recorder.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
	body := recorder.Body.String()
	assert.Equal(t, "OK", body)
}
