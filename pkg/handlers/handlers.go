package handlers

import (
	"WebApp/db/db_models"
	"WebApp/pkg/config"
	"WebApp/pkg/models"
	"WebApp/pkg/render"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Repo the repository used by the handlers
var (
	Repo *Repository
	// templateData = make(map[string]*models.TemplateData)
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

/*
SetupRoutes configures the HTTP route for a given template. The route is determined
based on the URL path of incoming requests. The last element of the URL path is used
as the key to retrieve the corresponding template from the PageTemplates map.

The function will also attempt to read the request body and prepare relevant template data.
If any step encounters an error, such as failing to read the request body, or if the desired
route isn't found in the PageTemplates map, an appropriate HTTP error response is sent back.

If everything goes as expected and the route is found in the map, the template is rendered
using the prepared template data.
*/

func (m *Repository) SetupRoutes() {
	lastElement := ""
	urlSeparator := "/"

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		templateData := models.TemplateData{
			StringMap: make(map[string]string),
		}
		body, err := io.ReadAll(r.Body)
		url := strings.Split(r.URL.Path, urlSeparator)
		lastElement = url[len(url)-1]
		if err != nil {
			// Handle the error here
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		err = json.Unmarshal(body, &templateData)
		if err != nil {
			http.Error(w, "Failed to prepare template data", http.StatusInternalServerError)
		}
		if len(body) != 0 {
			if templateData.DataType == "template" {
				val, exists := m.App.PageTemplates[lastElement]
				if exists {
					render.RenderTemplate(w, val, &templateData)
				} else {
					http.Error(w, fmt.Sprintf("Route %s is not in the map.", lastElement), http.StatusInternalServerError)
				}
			} else if templateData.DataType == "database" {
				db_models.CreateRecord(m.App.DbConnection, body)
				val, exists := m.App.PageTemplates[lastElement]
				if exists {
					render.RenderTemplate(w, val, nil)
				} else {
					http.Error(w, fmt.Sprintf("Route %s is not in the map.", lastElement), http.StatusInternalServerError)
				}
			}
		}
	}
	http.HandleFunc(urlSeparator+lastElement, handlerFunc)
}
