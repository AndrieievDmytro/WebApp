package handlers

import (
	"WebApp/pkg/config"
	"WebApp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var (
	Repo *Repository
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

func (m *Repository) RenderPage(w http.ResponseWriter, r *http.Request, pageTmplName string) {
	render.RenderTemplate(w, m.App.PageTemplates[pageTmplName])
}

func (m *Repository) SetupRoutes() {
	for key, val := range m.App.PageTemplates {
		// Capture the value in a closure
		handlerFunc := func(val string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				render.RenderTemplate(w, val)
			}
		}(val)
		http.HandleFunc("/"+key, handlerFunc)
	}
}
