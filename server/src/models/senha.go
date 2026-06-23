package models

// Representa o payload que recebe os valores para atualizar a senha do usuario
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
