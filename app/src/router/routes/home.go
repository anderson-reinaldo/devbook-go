package routes

import (
	"app/src/controllers"
	"net/http"
)

var routesHome = Rota{
	URI:      "/home",
	Method:   http.MethodGet,
	Function: controllers.LoadHomeScreen,
	IsAuth:   true,
}
