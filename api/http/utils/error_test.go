package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	Name  string `validate:"required"`
	Age   int    `validate:"gte=18"`
	Email string `validate:"required,email"`
}

func TestValidateBody_ValidBody(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	body := SampleStruct{
		Name:  "John Doe",
		Age:   25,
		Email: "john.doe@example.com",
	}

	err := ValidateBody(r, body)
	assert.Nil(t, err)
}

func TestValidateBody_InvalidBody(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	body := SampleStruct{
		Name:  "",
		Age:   16,
		Email: "invalid-email",
	}

	err := ValidateBody(r, body)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Cause, "Name")
	assert.Contains(t, err.Cause, "Age")
	assert.Contains(t, err.Cause, "Email")
}

func TestNewError(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	err := NewError(http.StatusForbidden, "Access denied", "Forbidden access", r)

	assert.Equal(t, http.StatusForbidden, err.Status)
	assert.Equal(t, "Access denied", err.Message)
	assert.Equal(t, "Forbidden access", err.Cause)
	assert.Equal(t, "/test", err.Path)
	_, parseErr := time.Parse(time.RFC3339, err.Timestamp)
	assert.Nil(t, parseErr)
}

func TestNewInternalServerError(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	err := NewInternalServerError(errors.New("internal error"), r)

	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "Erro interno.", err.Message)
	assert.Equal(t, "internal error", err.Cause)
	assert.Equal(t, "/test", err.Path)
	_, parseErr := time.Parse(time.RFC3339, err.Timestamp)
	assert.Nil(t, parseErr)
}

func TestNewNotFoundError(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	err := NewNotFoundError(errors.New("not found"), r)

	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.Equal(t, "O recurso não foi encontrado.", err.Message)
	assert.Equal(t, "not found", err.Cause)
	assert.Equal(t, "/test", err.Path)
	_, parseErr := time.Parse(time.RFC3339, err.Timestamp)
	assert.Nil(t, parseErr)
}

func TestNewBadRequestError(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	err := NewBadRequestError(errors.New("bad request"), r)

	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Requisição inválida.", err.Message)
	assert.Equal(t, "bad request", err.Cause)
	assert.Equal(t, "/test", err.Path)
	_, parseErr := time.Parse(time.RFC3339, err.Timestamp)
	assert.Nil(t, parseErr)
}
