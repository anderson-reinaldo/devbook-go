package middlewares

import (
	"devbook/src/auth"
	"devbook/src/responses"
	"fmt"
	"net/http"
)

// Logger imprime informações da requisição
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Proto)
		next(w, r)
	}
}

// Autenticar verifica se o usuário está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := auth.ValidarToken(r); err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
