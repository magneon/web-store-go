package routes

import (
	"net/http"
	ctr "web-store-go/infra/entrypoints/web/controllers"
)

func Routes() {
	http.HandleFunc("/", ctr.Index)
}
