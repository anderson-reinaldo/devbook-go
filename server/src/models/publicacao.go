package models

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicacao criada por um Publicacao
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"AutorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

// Preparar formata e valida o usuário
func (u *Publicacao) Preparar() error {
	if err := u.validar(); err != nil {
		return err
	}

	err := u.formatar()
	if err != nil {
		return err
	}

	return nil
}

func (u *Publicacao) validar() error {
	if u.Conteudo == "" {
		return errors.New("O campo conteudo é obrigatório")
	}
	if u.Titulo == "" {
		return errors.New("O campo titulo é obrigatório")
	}

	if len(u.Conteudo) > 300 {
		return errors.New("O campo conteudo nao pode ultrapassar os 300 caracteres.")
	}

	return nil
}

func (u *Publicacao) formatar() error {
	u.Titulo = strings.TrimSpace(u.Titulo)
	u.Conteudo = strings.TrimSpace(u.Conteudo)

	return nil
}
