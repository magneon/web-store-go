package repository

import (
	_ "github.com/lib/pq"

	"database/sql"
)

func Conectar() *sql.DB {
	conexao := "user=root dbname=alura_loja host=localhost password=toor sslmode=disable"

	db, erro := sql.Open("postgres", conexao)
	if erro != nil {
		panic(erro.Error())
	}
	return db
}
