package repositorios

import (
	"database/sql"
	"go-api/src/modelos"
)

// struct para receber o banco
type Usuarios struct {
	db *sql.DB
}

// cria novo repositorio de usuario
func NovoRepositodioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// [insere] um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}

// [insere] um usuario no banco de dados
