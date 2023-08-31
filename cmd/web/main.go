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
	app        config.AppConfig
	portNumber string = ":8080"
)

func init() {
	flag.StringVar(&portNumber, "p", portNumber, "Port to run the application on")
}

func main() {

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TempateCache = templateCache

	//if set to false update template on each request(development mode)
	//if set to true write templates to cache
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	flag.Parse()
	fmt.Printf("Starting application on port %s \n", portNumber)

	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
