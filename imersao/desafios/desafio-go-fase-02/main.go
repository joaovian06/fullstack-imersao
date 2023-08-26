package main

import (
	"net/http"

	// "route-api/api"

	"github.com/go-chi/chi/v5"
	"github.com/joaovian06/desafio-go/tree/main/imersao/desafios/desafio-go-fase-02/api"
)

func main() {
	api.InitDB("root:root@tcp(mysql:3306)/routes")

	r := chi.NewRouter()

	r.Post("/api/routes", api.CreateRouteHandler)
	r.Get("/api/routes", api.ListRoutesHandler)

	http.ListenAndServe(":8080", r)
}