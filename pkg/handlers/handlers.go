package handlers

import (
	"net/http"
	"server/pkg/config"
	"server/pkg/models"
	"server/pkg/renders"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(rw, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	renders.RenderTemplate(rw, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
