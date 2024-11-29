package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"ready-server/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Letsgo)

	log.Print("starting server")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("shutting server down")
}
