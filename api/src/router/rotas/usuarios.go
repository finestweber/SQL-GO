package rotas

import (
	//"go-api/src/controllers"
	"go-api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	// ROTA QUE CRIA UM USUARIO {METHODPOST}
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticação: false,
	},
	// 	ROTA QUE BUSCA TODOS OS USUARIOS {METHODGET}
	{
		URI:                "/usuarios/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticação: false,
	},
	// ROTA QUE BUSCA UM USARIO {METHODGET}
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticação: false,
	},
	// ATUALIZA USUARIO
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizaUsuario,
		RequerAutenticação: false,
	},
	//	ROTA QUE DELETE UM USUARIO {METHODDELETE}
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletaUsuario,
		RequerAutenticação: false,
	},
}
