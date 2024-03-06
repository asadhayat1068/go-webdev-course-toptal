package main

import (
	"net/http"

	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/config"
	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	// Routes
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
