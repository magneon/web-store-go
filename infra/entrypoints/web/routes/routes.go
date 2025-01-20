package routes

import (
	"net/http"
	ctr "web-store-go/infra/entrypoints/web/controllers"
)

func Routes() {
	http.HandleFunc("/", ctr.Index)
	http.HandleFunc("/new", ctr.New)
	http.HandleFunc("/insert", ctr.Insert)
	http.HandleFunc("/delete", ctr.Delete)
	http.HandleFunc("/edit", ctr.Edit)
	http.HandleFunc("/update", ctr.Update)
}
