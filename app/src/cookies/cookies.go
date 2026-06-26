package cookies

import (
	"app/src/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// SetupSecureCookie usa as variaveis de ambiente para configurar o SecureCookie
func SetupSecureCookie() {
	s = securecookie.New(config.Hashkey, config.Blockkey)
}

// Save registra as informações de autenticação
func Save(w http.ResponseWriter, ID, token string) error {

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	dataEncoded, erro := s.Encode("data", data)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    dataEncoded,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Read ler as informações do cookie de autenticação
func Read(r *http.Request) (map[string]string, error) {

	cookie, erro := r.Cookie("data")
	if erro != nil {
		return nil, erro
	}

	value := make(map[string]string)
	if erro = s.Decode("data", cookie.Value, &value); erro != nil {
		return nil, erro
	}

	return value, nil
}
