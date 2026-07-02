package controllers

import (
	"devbook/src/auth"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarPublicacao cria uma publicacao
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var publicacao models.Publicacao
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	publicacao.AutorID = usuarioID

	if err := json.Unmarshal([]byte(body), &publicacao); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = publicacao.Preparar()
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

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacoes busca todas as publicacoes dos usuarios que o usuario que fez a requisição segue e do propio usuario
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)

	publicacoes, err := repositorio.Buscar(usuarioID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publicacoes)
}

// BuscarPublicacoes busca todas as publicacoes de um usuario
func BuscarPublicacoesUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 32)
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

	repositorio := repositories.NovoRepositorioDePublicacoes(db)

	publicacoes, err := repositorio.BuscarPublicacoesUsuario(usuarioId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publicacoes)
}

// BuscarPublicacao busca uma publicacao de um usuario
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	fmt.Println(publicacaoId)

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)

	publicacao, err := repositorio.BuscarPorId(publicacaoId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, publicacao)
}

// AtualizarPublicacao atualiza uma publicacao
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var publicacao models.Publicacao
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal([]byte(body), &publicacao); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = publicacao.Preparar()
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

	repositorio := repositories.NovoRepositorioDePublicacoes(db)

	var publicacaoNoBanco models.Publicacao
	publicacaoNoBanco, err = repositorio.BuscarPorId(publicacaoId)
	if err != nil {
		responses.ERROR(w, http.StatusForbidden, err)
		return
	}

	if publicacaoNoBanco.AutorID != usuarioID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Impossivel editar uma publicação de outro usuario"))
		return
	}

	err = repositorio.Atualizar(publicacaoId, publicacao)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletarPublicacao deleta uma publicacao
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	var publicacaoNoBanco models.Publicacao

	publicacaoNoBanco, err = repositorio.BuscarPorId(publicacaoId)
	if err != nil {
		responses.ERROR(w, http.StatusForbidden, err)
		return
	}

	if publicacaoNoBanco.AutorID != usuarioID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Impossivel apagar uma publicação de outro usuario"))
		return
	}
	err = repositorio.Deletar(publicacaoId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// CurtirPublicacao adiciona uma curtida em uma publicacao
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	var params = mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 32)
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

	repositorio := repositories.NovoRepositorioDePublicacoes(db)

	err = repositorio.Curtir(usuarioID, publicacaoId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Usuario já curtiu essa publicação"))
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DescurtirPublicacao remove uma curtida em uma publicacao
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	var params = mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 32)
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

	repositorio := repositories.NovoRepositorioDePublicacoes(db)

	err = repositorio.Descurtir(usuarioID, publicacaoId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
