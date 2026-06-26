package router

import (
	"app/src/router/routes"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	return routes.InitRoutes(mux.NewRouter())
}
