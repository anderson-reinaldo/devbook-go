package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	IsAuth   bool
}

func InitRoutes(router *mux.Router) *mux.Router {
	routes := rotasLogin
	routes = append(routes, routesUsers...)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
