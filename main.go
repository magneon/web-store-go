package main

import (
	"fmt"
	"net/http"
	"text/template"
	produtos "web-store-go/produtos"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	produtos := []produtos.Produto{
		{Nome: "Camiseta", Descricao: "Bem bonita", Preco: 29, Quantidade: 10},
		{Nome: "Notebook", Descricao: "Muito rápido", Preco: 1999, Quantidade: 1},
		{Nome: "Celular", Descricao: "Zica rápido", Preco: 3999, Quantidade: 3},
	}
	fmt.Println(produtos)

	temp.ExecuteTemplate(writer, "index", produtos)
}
