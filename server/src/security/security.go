package security

import "golang.org/x/crypto/bcrypt"

// Hash cria um hash da senha
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara a senha com o hash
func VerificarSenha(hashedSenha, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}
