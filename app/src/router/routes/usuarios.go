package routes

import (
	"app/src/controllers"
	"net/http"
)

var routesUsers = []Rota{
	{URI: "/criar-usuario", Method: http.MethodGet, Function: controllers.LoadRegisterUserScreen, IsAuth: false},
	{URI: "/usuarios", Method: http.MethodPost, Function: controllers.CreateUser, IsAuth: false},
	{URI: "/usuarios", Method: http.MethodGet, Function: controllers.BuscarUsuarios, IsAuth: true},
	{URI: "/usuarios/{usuarioId}", Method: http.MethodPut, Function: controllers.AtualizarUsuario, IsAuth: true},
	{URI: "/usuarios/{usuarioId}", Method: http.MethodDelete, Function: controllers.DeletarUsuario, IsAuth: true},
	{URI: "/usuarios/{usuarioId}/seguindo", Method: http.MethodGet, Function: controllers.BuscarSeguindo, IsAuth: true},
	{URI: "/usuarios/{usuarioId}/seguir", Method: http.MethodPost, Function: controllers.SeguirUsuario, IsAuth: true},
	{URI: "/usuarios/{usuarioId}/parar-de-seguir", Method: http.MethodPost, Function: controllers.PararDeSeguirUsuario, IsAuth: true},
	{URI: "/usuarios/{usuarioId}/atualizar-senha", Method: http.MethodPost, Function: controllers.AtualizarSenha, IsAuth: true},
}
