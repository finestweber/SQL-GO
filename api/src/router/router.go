package router

import (
	"go-api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um Router coa as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
