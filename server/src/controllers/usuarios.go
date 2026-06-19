package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario cria um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal([]byte(body), &usuario); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = usuario.Preparar("cadastro")
	if err != nil {
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
	usuario.ID, err = repositorio.Criar(usuario)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	usuarios, err := repositorio.Buscar(nameOrNick)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	usuarios, err := repositorio.BuscarPorId(usuarioId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

// AtualizarUsuario atualiza um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	var usuario models.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal([]byte(body), &usuario); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = usuario.Preparar("atualizacao")
	if err != nil {
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

	err = repositorio.Atualizar(usuarioId, usuario)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario deleta um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	err = repositorio.Deletar(usuarioId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
