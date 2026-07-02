package routes

import (
	"app/src/controllers"
	"net/http"
)

var routesHome = []Rota{
	{URI: "/home", Method: http.MethodGet, Function: controllers.LoadHomeScreen, IsAuth: true},
	{URI: "/pagina/usuarios", Method: http.MethodGet, Function: controllers.LoadUsuariosScreen, IsAuth: true},
	{URI: "/perfil/{usuarioId}", Method: http.MethodGet, Function: controllers.LoadPerfilScreen, IsAuth: true},
}
