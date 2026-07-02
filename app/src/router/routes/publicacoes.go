package routes

import (
	"app/src/controllers"
	"net/http"
)

var routesPublicacoes = []Rota{
	{URI: "/publicacoes", Method: http.MethodPost, Function: controllers.CriarPublicacao, IsAuth: true},
	{URI: "/publicacoes/{publicacaoId}", Method: http.MethodPut, Function: controllers.AtualizarPublicacao, IsAuth: true},
	{URI: "/publicacoes/{publicacaoId}", Method: http.MethodDelete, Function: controllers.DeletarPublicacao, IsAuth: true},
	{URI: "/publicacoes/{publicacaoId}/curtir", Method: http.MethodPost, Function: controllers.CurtirPublicacao, IsAuth: true},
	{URI: "/publicacoes/{publicacaoId}/descurtir", Method: http.MethodPost, Function: controllers.DescurtirPublicacao, IsAuth: true},
}
