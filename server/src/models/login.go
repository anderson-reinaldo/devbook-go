package models

// Representa o payload devolve o token e o ID do usuario que efetuou login
type UsuarioToken struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
