package models

import "time"

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
