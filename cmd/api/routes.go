package main

import (
	"github.com/RealImage/challenge2016/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/v1/ping"))

	//mux.Get("/v1/dataset", service.LoadDatasetHandler)
	mux.Post("/v1/distribute/{name}", service.CreateDistributorHandler)
	mux.Post("/v1/sub-distribute/{parent_id}/{name}", service.CreateSubDistributorHandler)
	mux.Post("/v1/check-permissions/{id}", service.CheckDistributorPermissionHandler)

	return mux
}
