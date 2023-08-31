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
	portNumber string = ":8080"
)

func init() {
	flag.StringVar(&portNumber, "p", portNumber, "Port to run the application on")
}

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TempateCache = templateCache

	flag.Parse()
	fmt.Printf("Starting application on port %s \n", portNumber)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
