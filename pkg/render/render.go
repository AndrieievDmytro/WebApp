package render

import (
	"WebApp/pkg/config"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	app *config.AppConfig
)

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var templateCache map[string]*template.Template
	// Check if the app is in development mode or not
	if app.UseCache {
		// Retrieve template cache from app config
		templateCache = app.TempateCache
	} else {
		var err error
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Println("Error creating template cache:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	// Get the requested template
	templateSet, exists := templateCache[tmpl]
	if !exists {
		log.Println("Template not found in cache:", tmpl)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	buffer := new(bytes.Buffer)
	// Render the template into the buffer
	err := templateSet.Execute(buffer, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Write the buffer contents to the HTTP response writer
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// Create an empty template cache.
	myCache := make(map[string]*template.Template)
	app.PageTemplates = make(map[string]string)
	// Load layout templates only once since they won't change per page iteration.
	matches, err := filepath.Glob(app.LayoutTemlatesPath)
	if err != nil {
		return nil, fmt.Errorf("error getting layout templates: %w", err)
	}
	// Get all the page templates.
	pages, err := filepath.Glob(app.PageTemplatesPath)
	if err != nil {
		return nil, fmt.Errorf("error getting page templates: %w", err)
	}
	// Range through all page templates.
	for _, page := range pages {
		name := filepath.Base(page)
		//Populate the PageTemplates dictionary in AppConfig structure using the prefix (first word) of each template filename as the key.
		shortName := strings.Split(name, ".")
		if len(shortName) > 0 {
			app.PageTemplates[shortName[0]] = name
		}
		// Parse the page template.
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, fmt.Errorf("error parsing page template %s: %w", name, err)
		}
		// If there are layout templates, parse them.
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(app.LayoutTemlatesPath)
			if err != nil {
				return nil, fmt.Errorf("error parsing layout templates for page %s: %w", name, err)
			}
		}
		// Store the complete template set in the cache.
		myCache[name] = ts
	}
	return myCache, nil
}
