package config

import (
	"log"
	"text/template"
)

// AppConfig contains the application configurations
type AppConfig struct {
	UseCache       bool
	TempateCache   map[string]*template.Template
	InfoLog        *log.Logger
	PageTemplates  string
	LayoutTemlates string
	PortNumber     string
}

// Configure the basic configurations of the application
func (a *AppConfig) ConfigureApp(isProdMode bool) {
	a.PageTemplates = "C:\\ProjectFolder\\Go\\src\\WebApp\\templates\\*.page.tmpl"
	a.LayoutTemlates = "C:\\ProjectFolder\\Go\\src\\WebApp\\templates\\*.layout.tmpl"
	a.PortNumber = ":8080"
	a.UseCache = isProdMode
}
