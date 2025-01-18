package main

import (
	"database/sql"
	prd "web-store-go/produtos"

	_ "github.com/lib/pq"

	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func conectar() *sql.DB {
	conexao := "user=root dbname=alura_loja host=localhost password=toor sslmode=disable"

	db, erro := sql.Open("postgres", conexao)
	if erro != nil {
		panic(erro.Error())
	}
	return db
}

func index(writer http.ResponseWriter, request *http.Request) {
	db := conectar()
	resultado, erro := db.Query("select * from produtos")
	if erro != nil {
		panic(erro.Error())
	}

	produtos := []prd.Produto{}
	produto := prd.Produto{}
	for resultado.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err := resultado.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
		produtos = append(produtos, produto)
	}

	temp.ExecuteTemplate(writer, "index", produtos)
	defer db.Close()
}
