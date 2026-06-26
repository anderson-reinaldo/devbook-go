package routes

import (
	"app/src/controllers"
	"net/http"
)

var routesLogin = []Rota{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Function: controllers.LoadLoginScreen,
		IsAuth:   false,
	},
	{
		URI:      "/login",
		Method:   http.MethodGet,
		Function: controllers.LoadLoginScreen,
		IsAuth:   false,
	},
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Function: controllers.HandleLogin,
		IsAuth:   false,
	},
}
