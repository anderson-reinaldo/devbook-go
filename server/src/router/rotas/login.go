package rotas

import (
	"devbook/src/controllers"
	"net/http"
)

var rotasLogin = []Rota{
	{
		URI:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Login,
		RequerAutenticacao: false,
	},
}
