package config

import (
	"log"
	"text/template"
)

// AppConfig contains the application configuration
type AppConfig struct {
	UseCache     bool
	TempateCache map[string]*template.Template
	InfoLog      *log.Logger
}
