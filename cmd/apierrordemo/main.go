package main

import (
	"log"
	"net/http"

	"github.com/alesr/apierrordemo/internal/api"
	"github.com/alesr/apierrordemo/internal/repository"
	"github.com/alesr/apierrordemo/internal/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	repo := repository.NewFakePostgres()
	service := service.NewDefaultService(repo)

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("failed to initialize custom logger:", err)
	}

	app := api.New(logger, service)

	r := chi.NewRouter()

	r.Get("/users/{id}", app.FetchUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln("failed to serve on :8080:", err)
	}
}
