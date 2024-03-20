package main

import (
	"net/http"

	"github.com/dreamsofcode-io/nethttp/monster"
)

func loadRoutes(router *http.ServeMux) {
	handler := &monster.Handler{}

	router.HandleFunc("POST /monster", handler.Create)
	router.HandleFunc("PUT /monster/{id}", handler.UpdateByID)
	router.HandleFunc("GET /monster/{id}", handler.FindByID)
	router.HandleFunc("DELETE /monster/{id}", handler.DeleteByID)
}
