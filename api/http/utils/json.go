package utils

import (
	"encoding/json"
	"net/http"
)

// Função genérica responsável por criar uma resposta JSON.
// Recebe um http.ResponseWriter, um status HTTP e um objeto que será serializado.
func NewJSONResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
