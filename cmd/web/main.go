package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/config"
	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/handlers"
	"github.com/asadhayat1068/gowebdev_toptal_udemy/pkg/render"
)

const PORT = ":8080"

func main() {
	// Init App Configs
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.UseCache = false
	app.TemplateCache = tc

	// Init Handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Init render
	render.NewTemplates(&app)

	// Routes
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Server is running at port", PORT)
	// http.ListenAndServe(PORT, nil)

	srv := http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
