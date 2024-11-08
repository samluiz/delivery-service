package server

import (
	"database/sql"
	"net/http"
)

// Struct que representa o servidor HTTP.
// Contém um mux e um banco de dados como atributos.
type Server struct {
	db     *sql.DB        // Banco de dados
	Router *http.ServeMux // Mux (router)
}

// Função responsável por instanciar um novo server.
func NewServer(db *sql.DB) *Server {
	return &Server{
		db:     db,
		Router: http.NewServeMux(),
	}
}

// Delegando o método ServeHTTP para o request multiplexer (router).
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
