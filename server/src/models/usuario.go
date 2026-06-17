package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário
type Usuario struct {
	ID       uint64    `json:id,omitempty`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar formata e valida o usuário
func (u *Usuario) Preparar() error {
	if err := u.validar(); err != nil {
		return err
	}
	u.formatar()

	return nil
}

func (u *Usuario) validar() error {

	if u.Nome == "" {
		return errors.New("O campo nome é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("O campo nick é obrigatório")
	}
	if u.Email == "" {
		return errors.New("O campo email é obrigatório")
	}
	if u.Senha == "" {
		return errors.New("O campo senha é obrigatório")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
