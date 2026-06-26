package routes

import (
	"app/src/controllers"
	"net/http"
)

var routesUsers = []Rota{
	{
		URI:      "/criar-usuario",
		Method:   http.MethodGet,
		Function: controllers.LoadRegisterUserScreen,
		IsAuth:   false,
	},
	{
		URI:      "/usuarios",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		IsAuth:   false,
	},
}
