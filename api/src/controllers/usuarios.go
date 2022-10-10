package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/src/banco"
	"go-api/src/modelos"
	"go-api/src/repositorios"
	"io/ioutil"
	"log"
	"net/http"
)

// criar usuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("cria usuario no banco"))
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)

	}
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositodioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("id inserido: %d", usuarioID)))
}

// BUSCA TODOS OS USUARIOS NO BANCO
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos os usuarios"))
}

// BUSCA UM USUARIO NO BANCO
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca um usuario"))
}

// ATUALIZA USUARIO NO BANCO
func AtualizaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

// DELETA USUARIO NO BANCO
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario"))
}
