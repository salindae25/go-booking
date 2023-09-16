package handlers

import (
	"net/http"

	"github.com/salindae25/go-booking/pkg/config"
	"github.com/salindae25/go-booking/pkg/models"
	"github.com/salindae25/go-booking/pkg/renders"
)

// Repo repository used by handler
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo create new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set Repo with repository
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteAddr)
	renders.RenderTemplate(w, "home", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Name"] = "Salinda"

	remoteAddr := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteAddr

	renders.RenderTemplate(w, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}
