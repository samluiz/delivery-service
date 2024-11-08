package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Validator para validação de erros.
var v = &XValidator{
	Validator: validator.New(),
}

// Struct que representa um erro HTTP.
type Error struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Cause     string `json:"error"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path"`
}

// Função responsável por validar um struct utilizando o validator.
func ValidateBody(r *http.Request, body interface{}) *Error {

	// Validando o struct e retornando os erros se houver
	if errs := v.Validate(body); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		// Iterando sobre os erros e adicionando a mensagem de erro para o usuário
		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Deve satisfazer a validação '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		// Concatenando os erros para serem exibidos para o usuário
		validationError := errors.New(strings.Join(errMsgs, " e "))

		// Retornando o erro de validação
		return NewBadRequestError(validationError, r)
	}
	return nil
}

// Função responsável por criar um erro HTTP.
func NewError(status int, message string, cause string, r *http.Request) *Error {
	return &Error{
		Status:    status,
		Message:   message,
		Cause:     cause,
		Timestamp: time.Now().Format(time.RFC3339),
		Path:      r.URL.Path,
	}
}

// Função responsável por criar um erro de servidor interno.
func NewInternalServerError(err error, r *http.Request) *Error {
	return &Error{
		Status:    http.StatusInternalServerError,
		Message:   "Erro interno.",
		Cause:     err.Error(),
		Timestamp: time.Now().Format(time.RFC3339),
		Path:      r.URL.Path,
	}
}

// Função responsável por criar um erro de recurso não encontrado.
func NewNotFoundError(err error, r *http.Request) *Error {
	return &Error{
		Status:    http.StatusNotFound,
		Message:   "O recurso não foi encontrado.",
		Cause:     err.Error(),
		Timestamp: time.Now().Format(time.RFC3339),
		Path:      r.URL.Path,
	}
}

// Função responsável por criar um erro de requisição inválida.
func NewBadRequestError(err error, r *http.Request) *Error {
	return &Error{
		Status:    http.StatusBadRequest,
		Message:   "Requisição inválida.",
		Cause:     err.Error(),
		Timestamp: time.Now().Format(time.RFC3339),
		Path:      r.URL.Path,
	}
}
