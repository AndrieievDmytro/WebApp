package config

import (
	"log"
	"text/template"

	"gorm.io/gorm"
)

// AppConfig structure contains the application configurations
type AppConfig struct {
	UseCache           bool
	TempateCache       map[string]*template.Template
	InfoLog            *log.Logger
	PageTemplatesPath  string
	LayoutTemlatesPath string
	PortNumber         string
	PageTemplates      map[string]string
	Dsn                string
	DbConnection       *gorm.DB
}

// Configure the basic configurations of the application
func (a *AppConfig) ConfigureApp(isProdMode bool) {
	a.PageTemplatesPath = "C:\\ProjectFolder\\Go\\src\\WebApp\\templates\\*.page.tmpl"
	a.LayoutTemlatesPath = "C:\\ProjectFolder\\Go\\src\\WebApp\\templates\\*.layout.tmpl"
	a.PortNumber = ":8080"
	a.UseCache = isProdMode
	a.Dsn = "host=localhost user=postgres password=test123 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
}
