package controllers

import (
	"devbook/src/auth"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"devbook/src/security"
	"encoding/json"
	"errors"
	"fmt"
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
		return
	}

	var usuario models.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	usuarioIdToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioId != usuarioIdToken {
		responses.ERROR(w, http.StatusForbidden, fmt.Errorf("Não é possível atualizar um usuário que não seja o seu"))
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
		return
	}

	usuarioIdToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioId != usuarioIdToken {
		responses.ERROR(w, http.StatusForbidden, fmt.Errorf("Não é possível deletar um usuário que não seja o seu"))
		return
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

// SeguirUsuario segue um usuario
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if seguidorId == usuarioId {
		responses.ERROR(w, http.StatusForbidden, fmt.Errorf("Não é possível seguir você mesmo"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	err = repositorio.Seguir(usuarioId, seguidorId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario para de seguir um usuario
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if seguidorId == usuarioId {
		responses.ERROR(w, http.StatusForbidden, fmt.Errorf("Não é possível você deixar de seguir você mesmo"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	err = repositorio.PararDeSeguir(usuarioId, seguidorId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores busca os seguidores de um usuario
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	seguidores, err := repositorio.Seguidores(usuarioId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo busca os quem um usuario está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	seguidores, err := repositorio.Seguindo(usuarioId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo busca os quem um usuario está seguindo
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	usuarioIdToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioId != usuarioIdToken {
		responses.ERROR(w, http.StatusForbidden, fmt.Errorf("Não é possível atualizar a senha de outro usuario"))
		return
	}

	var senha models.Senha
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal([]byte(body), &senha); err != nil {
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
	senhaDoBanco, err := repositorio.BuscarSenha(usuarioId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = security.VerificarSenha(senhaDoBanco, senha.Atual)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Senha Atual informada diferente da salva."))
		return
	}

	novaSenhaHash, err := security.Hash(senha.Nova)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = repositorio.AtualizarSenha(usuarioId, string(novaSenhaHash))
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
