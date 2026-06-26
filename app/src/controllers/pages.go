package controllers

import (
	"app/src/config"
	"app/src/models"
	"app/src/requests"
	"app/src/response"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// LoadLoginScreen carrega a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecutorTemplate(w, "login.html", nil)
}

// LoadLoginScreen carrega a tela de login
func LoadRegisterUserScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecutorTemplate(w, "register.html", nil)
}

// LoadHomeScreen carrega a pagina inicial com as publicacoes
func LoadHomeScreen(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.BASEURL_API)
	fmt.Println(url)
	res, err := requests.HandleRequestAuth(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	var publicacoes []models.Publicacao

	if err := json.NewDecoder(res.Body).Decode(&publicacoes); err != nil {
		response.JSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.ExecutorTemplate(w, "home.html", publicacoes)
}
