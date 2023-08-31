package render

import (
	"WebApp/pkg/config"
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	pageTemplates  = "./templates/*.page.tmpl"
	layoutTemlates = "./templates/*.layout.tmpl"
	app            *config.AppConfig
)

// NewTemplates sets the config for the template function
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	//if/else statment checks if the app in development mode or not
	if app.UseCache {
		//recieving template cache from app config
		tc = app.TempateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	// Check if something wrong with templates stored in the map, do not execute directly.
	_ = t.Execute(buf, nil)

	//render the tempate
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files end with *.page.tmpl from ./templates

	pages, err := filepath.Glob(pageTemplates)
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		//populate template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(layoutTemlates)
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(layoutTemlates)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
