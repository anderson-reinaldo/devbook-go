package auth

import (
	"devbook/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GerarToken gera um token com as permissões do usuário
func GerarToken(id uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado %v", token.Header["alg"])
	}

	return []byte(config.JwtSecret), nil
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, err := strconv.ParseUint(fmt.Sprintf("%.f", permissions["id"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}
