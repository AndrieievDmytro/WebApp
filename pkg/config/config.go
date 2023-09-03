package config

import (
	"log"
	"text/template"
)

// AppConfig contains the application configurations
type AppConfig struct {
	UseCache           bool
	TempateCache       map[string]*template.Template
	InfoLog            *log.Logger
	PageTemplatesPath  string
	LayoutTemlatesPath string
	PortNumber         string
	PageTemplates      map[string]string
}

// Configure the basic configurations of the application
func (a *AppConfig) ConfigureApp(isProdMode bool) {
	a.PageTemplatesPath = "C:\\ProjectFolder\\Go\\src\\WebApp\\templates\\*.page.tmpl"
	a.LayoutTemlatesPath = "C:\\ProjectFolder\\Go\\src\\WebApp\\templates\\*.layout.tmpl"
	a.PortNumber = ":8080"
	a.UseCache = isProdMode
}
