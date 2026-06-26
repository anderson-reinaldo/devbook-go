package main

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/router"
	"app/utils"
	"fmt"
	"log"
	"net/http"
)

func init() {
	err := config.Loader()
	if err != nil {
		panic(err)
	}
	utils.LoaderTemplates()
	cookies.SetupSecureCookie()
}

func main() {
	fmt.Printf("Rodando APP na porta %d!", config.Port)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
