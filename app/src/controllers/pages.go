package controllers

import (
	"app/utils"
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
