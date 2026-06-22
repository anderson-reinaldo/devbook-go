package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GerarToken64() string {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(chave)
}
