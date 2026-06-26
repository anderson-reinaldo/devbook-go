package controllers

import (
	"app/src/config"
	"app/src/response"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// LoadLoginScreen carrega a tela de login
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
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

	response.JSON(w, res.StatusCode, nil)
}
