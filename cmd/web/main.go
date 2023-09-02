package main

import (
	"WebApp/pkg/config"
	"WebApp/pkg/handlers"
	"WebApp/pkg/render"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	app config.AppConfig
	//if set to false update template on each request(development mode)
	//if set to true write templates to cache
	isProdMode bool
)

func init() {
	flag.BoolVar(&isProdMode, "pd", false, "If true production mode is on, if false production mode is off.")
}

func main() {
	flag.Parse()
	fmt.Printf("Production Mode: %v\n", isProdMode)
	app.ConfigureApp(isProdMode)
	render.NewTemplates(&app)
	//Initialize structure cache field with a templates cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TempateCache = templateCache
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	fmt.Printf("Starting application on port %s \n", app.PortNumber)
	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	_ = http.ListenAndServe(app.PortNumber, nil)
}
