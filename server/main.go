package main

import (
	"devbook/src/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.Carregar()
}

func main() {
	router := router.Gerar()
	fmt.Printf("Server is on port %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
