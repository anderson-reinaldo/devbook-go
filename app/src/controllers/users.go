package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/models"
	"app/src/requests"
	"app/src/response"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	senha := r.FormValue("senha")

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": email,
		"senha": senha,
	})
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.Erro{Erro: erro.Error()})
		return
	}

	res, erro := http.Post(fmt.Sprintf("%s/usuarios", config.BASEURL_API), "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: erro.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	credenciais, erro := json.Marshal(map[string]string{"email": email, "senha": senha})
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: erro.Error()})
		return
	}

	resLogin, erro := http.Post(fmt.Sprintf("%s/login", config.BASEURL_API), "application/json", bytes.NewBuffer(credenciais))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: erro.Error()})
		return
	}
	defer resLogin.Body.Close()

	if resLogin.StatusCode >= 400 {
		response.HandleStatusCode(w, resLogin)
		return
	}

	var dataAuth models.UsuarioToken
	if erro = json.NewDecoder(resLogin.Body).Decode(&dataAuth); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Erro{Erro: erro.Error()})
		return
	}

	if erro = cookies.Save(w, dataAuth.ID, dataAuth.Token); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Erro{Erro: erro.Error()})
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	usuario := r.URL.Query().Get("usuario")
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.BASEURL_API, usuario)
	res, err := requests.HandleRequestAuth(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/usuarios/%s/seguir", config.BASEURL_API, vars["usuarioId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPost, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/usuarios/%s/parar-de-seguir", config.BASEURL_API, vars["usuarioId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPost, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/usuarios/%s", config.BASEURL_API, vars["usuarioId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPut, url, r.Body)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/usuarios/%s", config.BASEURL_API, vars["usuarioId"])
	res, err := requests.HandleRequestAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/usuarios/%s/seguindo", config.BASEURL_API, vars["usuarioId"])
	res, err := requests.HandleRequestAuth(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf("%s/usuarios/%s/atualizar-senha", config.BASEURL_API, vars["usuarioId"])
	res, err := requests.HandleRequestAuth(r, http.MethodPost, url, r.Body)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	response.Pipe(w, res)
}
