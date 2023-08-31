package main

import (
	"WebApp/pkg/handlers"
	"flag"
	"fmt"
	"net/http"
)

var (
	portNumber string = ":8080"
)

func init() {
	flag.StringVar(&portNumber, "p", portNumber, "Port to run the application on")
}

func main() {
	flag.Parse()
	fmt.Printf("Starting application on port %s \n", portNumber)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
