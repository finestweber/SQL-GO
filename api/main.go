package main

import (
	"fmt"
	"go-api/src/config"
	"go-api/src/router"
	"log"
	"net/http"
)

func main() {
	config.Carregando()
	fmt.Println("Rodando Api porta : 5000")

	//fmt.Println(config.StringConexaoBanco)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
