package routes

import (
	"app/src/middlewares"
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
	routes := routesLogin
	routes = append(routes, routesUsers...)
	routes = append(routes, routesHome)

	for _, route := range routes {
		if route.IsAuth {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticated(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
