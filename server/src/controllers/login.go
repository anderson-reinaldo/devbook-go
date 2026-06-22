package controllers

import (
	"devbook/src/auth"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"devbook/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login realiza o login
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal([]byte(body), &credentials); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioPorEmail, err := repositorio.BuscarPorEmail(credentials.Email)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuarioPorEmail.Senha, credentials.Senha); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GerarToken(usuarioPorEmail.ID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)

}
