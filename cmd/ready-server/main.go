package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"

	"ready-server/internal/config"
	"ready-server/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Letsgo)

	path := cfg.Host + ":" + cfg.Port

	log.Print("starting server at %s", path)
	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("shutting server down")
}
