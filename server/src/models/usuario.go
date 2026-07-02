package models

import (
	"devbook/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário
type Usuario struct {
	ID       int64     `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar formata e valida o usuário
func (u *Usuario) Preparar(etapa string) error {
	if err := u.validar(etapa); err != nil {
		return err
	}

	err := u.formatar(etapa)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usuario) validar(etapa string) error {

	if u.Nome == "" {
		return errors.New("O campo nome é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("O campo nick é obrigatório")
	}
	if u.Email == "" {
		return errors.New("O campo email é obrigatório")
	}

	if checkmail.ValidateFormat(u.Email) != nil {
		return errors.New("O email é inválido")
	}

	if u.Senha == "" && etapa == "cadastro" {
		return errors.New("O campo senha é obrigatório")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, err := security.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(senhaComHash)
	}

	return nil
}
