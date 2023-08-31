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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
