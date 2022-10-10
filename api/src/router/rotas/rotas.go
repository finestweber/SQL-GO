package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)
//TODAS AS ROTAS DO API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticação bool
}
// CONFIGURAR COLOCA TODAS AS ROTAS NO ROUTER
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
