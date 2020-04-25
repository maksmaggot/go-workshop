package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"workshop/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("start server")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("shutting down server")
}
