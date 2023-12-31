package main

import (
	"WebApp/db"
	"WebApp/db/db_models"
	"WebApp/helper"
	"WebApp/pkg/config"
	"WebApp/pkg/handlers"
	"WebApp/pkg/render"
	"flag"
	"fmt"
	"net/http"
)

var (
	app        config.AppConfig
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
	helper.CheckErr(err)
	app.TempateCache = templateCache
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	dbConnection := db.InitDb(&app)
	app.DbConnection = dbConnection
	db_models.InitModels(dbConnection)
	fmt.Printf("Starting application on port %s \n", app.PortNumber)
	repo.SetupRoutes()
	// repo.HandleRoutes()
	_ = http.ListenAndServe(app.PortNumber, nil)
}
