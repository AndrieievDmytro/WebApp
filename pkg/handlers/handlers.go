package handlers

import (
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
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			// Handle the error here
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		url := strings.Split(r.URL.Path, "/")
		lastElement = url[len(url)-1]
		templateData, err := prepareTemplateData(lastElement, body)
		if err != nil {
			// Handle the error here
			http.Error(w, "Failed to prepare template data", http.StatusInternalServerError)
			return
		}

		val, exists := m.App.PageTemplates[lastElement]
		if exists {
			render.RenderTemplate(w, val, templateData)
		} else {
			http.Error(w, fmt.Sprintf("Route %s is not in the map.", lastElement), http.StatusInternalServerError)
		}
	}
	http.HandleFunc("/"+lastElement, handlerFunc)
}

/*
prepareTemplateData takes a key string and a byte slice representing the body of an HTTP request.
It attempts to unmarshal the body into a TemplateData struct, specifically targeting the StringMap field.

This function is primarily used to extract structured data from the incoming request payload (in JSON format)
and prepare it for rendering within templates.

Parameters:
- key: Not directly used in the function body but could potentially serve a purpose in future implementations or in other parts of the codebase.
- body: A byte slice containing the JSON payload from an HTTP request, expected to match the structure of the TemplateData struct.

Returns:
- A pointer to a TemplateData instance populated with the unmarshaled data.
- An error, if the unmarshaling process fails. Otherwise, returns nil.
*/

func prepareTemplateData(key string, body []byte) (*models.TemplateData, error) {
	td := models.TemplateData{
		StringMap: make(map[string]string),
	}
	err := json.Unmarshal(body, &td)
	if err != nil {
		return &td, err
	}
	return &td, nil
}
