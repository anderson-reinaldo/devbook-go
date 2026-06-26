package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/models"
	"app/src/response"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.Erro{Erro: erro.Error()})
		return
	}

	res, erro := http.Post(fmt.Sprintf("%s/login", config.BASEURL_API), "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: erro.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	var dataAuth models.UsuarioToken
	if erro = json.NewDecoder(res.Body).Decode(&dataAuth); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Erro{Erro: erro.Error()})
		return
	}

	erro = cookies.Save(w, dataAuth.ID, dataAuth.Token)
	if erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Erro{Erro: erro.Error()})
		return
	}

	response.JSON(w, http.StatusOK, nil)

}
