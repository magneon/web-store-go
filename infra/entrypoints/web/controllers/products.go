package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	prd "web-store-go/application/domain/products"
)

var temp = template.Must(template.ParseGlob("infra/templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	products := prd.GetAllProducts()
	temp.ExecuteTemplate(writer, "index", products)
}

func New(writer http.ResponseWriter, request *http.Request) {
	temp.ExecuteTemplate(writer, "new", nil)
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		name := request.FormValue("name")
		description := request.FormValue("description")

		quantity, error := strconv.Atoi(request.FormValue("quantity"))
		if error != nil {
			log.Println("Falha ao converter a quantidade", error)
		}

		price, error := strconv.ParseFloat(request.FormValue("price"), 64)
		if error != nil {
			log.Println("Falha ao converter o preço", error)
		}

		prd.CreateProduct(name, description, quantity, price)
	}
	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	prd.DeleteProduct(id)
	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	id, error := strconv.Atoi(request.URL.Query().Get("id"))
	if error != nil {
		panic(error.Error())
	}
	product := prd.GetProduct(id)
	temp.ExecuteTemplate(writer, "edit", product)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id, error := strconv.Atoi(request.FormValue("id"))
		if error != nil {
			panic(error.Error())
		}

		price, error := strconv.ParseFloat(request.FormValue("price"), 64)
		if error != nil {
			panic(error.Error())
		}

		quantity, error := strconv.Atoi(request.FormValue("quantity"))
		if error != nil {
			panic(error.Error())
		}

		name := request.FormValue("name")
		description := request.FormValue("description")

		prd.UpdateProduct(id, name, description, price, quantity)
		http.Redirect(writer, request, "/", http.StatusMovedPermanently)
	}
}
