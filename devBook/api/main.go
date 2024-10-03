package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("rodando api")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintln(":%d", config.Porta), r))
}
