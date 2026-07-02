package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Port é onde vai ser armazenada a porta que o servidor web vai rodar
	Port int
	//BASEURL_API é o a string base de conexão com a API
	BASEURL_API string
	//HashKey é usada para autenticar o cookie
	Hashkey []byte
	//HashKey é usada para criptografar os dados do cookie
	Blockkey []byte
)

func Loader() error {
	var err error
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 5000
	}

	BASEURL_API = os.Getenv("BASEURL_API")
	Hashkey = []byte(os.Getenv("HASH_KEY"))
	Blockkey = []byte(os.Getenv("BLOCK_KEY"))

	return nil
}
