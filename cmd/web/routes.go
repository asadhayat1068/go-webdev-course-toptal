package main

import (
	"net/http"

	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/config"
	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/handlers"
	"github.com/go-chi/chi/v4"
)

func routes(app *config.AppConfig) http.Handler {
	// // Using pat
	// mux := pat.New()
	// // Routes
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// Using chi
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
