package controllers

import (
	"app/src/config"
	"app/src/requests"
	"app/src/response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.BASEURL_API)
	res, err := requests.HandleRequestAuth(r, http.MethodPost, url, r.Body)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/publicacoes/%s", config.BASEURL_API, vars["publicacaoId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPut, url, r.Body)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/publicacoes/%s", config.BASEURL_API, vars["publicacaoId"])
	res, err := requests.HandleRequestAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/publicacoes/%s/curtir", config.BASEURL_API, vars["publicacaoId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPost, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/publicacoes/%s/descurtir", config.BASEURL_API, vars["publicacaoId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPost, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}
