package rotas

import (
	"devbook/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLogin...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
			continue
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
			continue
		}
	}

	return r
}
