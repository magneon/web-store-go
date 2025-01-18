package controllers

import (
	"html/template"
	"net/http"

	prd "web-store-go/application/domain/products"
)

var temp = template.Must(template.ParseGlob("infra/templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	products := prd.GetAllProducts()
	temp.ExecuteTemplate(writer, "index", products)
}
