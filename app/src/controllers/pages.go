package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/models"
	"app/src/requests"
	"app/src/response"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecutorTemplate(w, "login.html", nil)
}

func LoadRegisterUserScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecutorTemplate(w, "register.html", nil)
}

func LoadHomeScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	idLogado, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// Busca nick do usuário logado
	var nickLogado string
	resUsuario, err := requests.HandleRequestAuth(r, http.MethodGet, fmt.Sprintf("%s/usuarios/%d", config.BASEURL_API, idLogado), nil)
	if err == nil && resUsuario.StatusCode < 400 {
		var usuario models.Usuario
		json.NewDecoder(resUsuario.Body).Decode(&usuario)
		resUsuario.Body.Close()
		nickLogado = usuario.Nick
	}

	res, err := requests.HandleRequestAuth(r, http.MethodGet, fmt.Sprintf("%s/publicacoes", config.BASEURL_API), nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	var publicacoes []models.Publicacao
	if err := json.NewDecoder(res.Body).Decode(&publicacoes); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Erro{Erro: err.Error()})
		return
	}

	utils.ExecutorTemplate(w, "home.html", models.DadosHome{
		Publicacoes:       publicacoes,
		IDUsuarioLogado:   idLogado,
		NickUsuarioLogado: nickLogado,
	})
}

func LoadUsuariosScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	idLogado, _ := strconv.ParseUint(cookie["id"], 10, 64)
	utils.ExecutorTemplate(w, "usuarios.html", idLogado)
}

type resultPerfil struct {
	usuario    models.Usuario
	publicacoes []models.Publicacao
	seguidores []models.Usuario
	seguindo   []models.Usuario
	errUsuario error
}

func fetchPerfil(r *http.Request, baseURL, usuarioId string) resultPerfil {
	type chUsuario    struct{ v models.Usuario;           err error }
	type chPubs       struct{ v []models.Publicacao;      err error }
	type chSeguidores struct{ v []models.Usuario;         err error }
	type chSeguindo   struct{ v []models.Usuario;         err error }

	chU  := make(chan chUsuario, 1)
	chP  := make(chan chPubs, 1)
	chSg := make(chan chSeguidores, 1)
	chSn := make(chan chSeguindo, 1)

	go func() {
		res, err := requests.HandleRequestAuth(r, http.MethodGet, fmt.Sprintf("%s/usuarios/%s", baseURL, usuarioId), nil)
		if err != nil || res.StatusCode >= 400 { chU <- chUsuario{err: fmt.Errorf("not found")}; return }
		defer res.Body.Close()
		var u models.Usuario
		json.NewDecoder(res.Body).Decode(&u)
		chU <- chUsuario{v: u}
	}()

	go func() {
		res, err := requests.HandleRequestAuth(r, http.MethodGet, fmt.Sprintf("%s/usuarios/%s/publicacoes", baseURL, usuarioId), nil)
		if err != nil || res.StatusCode >= 400 { chP <- chPubs{}; return }
		defer res.Body.Close()
		var p []models.Publicacao
		json.NewDecoder(res.Body).Decode(&p)
		chP <- chPubs{v: p}
	}()

	go func() {
		res, err := requests.HandleRequestAuth(r, http.MethodGet, fmt.Sprintf("%s/usuarios/%s/seguidores", baseURL, usuarioId), nil)
		if err != nil || res.StatusCode >= 400 { chSg <- chSeguidores{}; return }
		defer res.Body.Close()
		var s []models.Usuario
		json.NewDecoder(res.Body).Decode(&s)
		chSg <- chSeguidores{v: s}
	}()

	go func() {
		res, err := requests.HandleRequestAuth(r, http.MethodGet, fmt.Sprintf("%s/usuarios/%s/seguindo", baseURL, usuarioId), nil)
		if err != nil || res.StatusCode >= 400 { chSn <- chSeguindo{}; return }
		defer res.Body.Close()
		var s []models.Usuario
		json.NewDecoder(res.Body).Decode(&s)
		chSn <- chSeguindo{v: s}
	}()

	rU  := <-chU
	rP  := <-chP
	rSg := <-chSg
	rSn := <-chSn

	return resultPerfil{
		usuario:     rU.v,
		publicacoes: rP.v,
		seguidores:  rSg.v,
		seguindo:    rSn.v,
		errUsuario:  rU.err,
	}
}

func LoadPerfilScreen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioId := vars["usuarioId"]

	cookie, _ := cookies.Read(r)
	idLogado, _ := strconv.ParseUint(cookie["id"], 10, 64)

	result := fetchPerfil(r, config.BASEURL_API, usuarioId)
	if result.errUsuario != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	seguindo := false
	for _, s := range result.seguidores {
		if s.ID == idLogado {
			seguindo = true
			break
		}
	}

	usuarioIdUint, _ := strconv.ParseUint(usuarioId, 10, 64)

	utils.ExecutorTemplate(w, "perfil.html", models.DadosPerfil{
		Usuario:         result.usuario,
		Publicacoes:     result.publicacoes,
		QtdSeguidores:   len(result.seguidores),
		QtdSeguindo:     len(result.seguindo),
		IDUsuarioLogado: idLogado,
		Seguindo:        seguindo,
		EhDono:          idLogado == usuarioIdUint,
	})
}
