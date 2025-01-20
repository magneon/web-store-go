package main

import (
	"net/http"
	"web-store-go/infra/entrypoints/web/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}
