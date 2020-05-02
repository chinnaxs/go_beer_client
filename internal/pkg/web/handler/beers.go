package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/chinnaxs/go_beer_client/internal/pkg/beverage"
)

var beersTemplateFileName = "beers.html"

var beersHandler *Handler

func BeersHandler(templatePath string, apiClient *api.ApiClient) http.HandlerFunc {

	beersTemplatePath := filepath.Join(templatePath, beersTemplateFileName)
	beersHandler = &Handler{
		Templates: template.Must(template.ParseFiles(
			beersTemplatePath,
		)),
		ApiClient: apiClient,
	}
	return beersHandler.listBeersHandler
}

func (h *Handler) listBeersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	if r.Method == http.MethodGet {
		beersHandler.getBeersHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (h *Handler) getBeersHandler(w http.ResponseWriter, r *http.Request) {
	beers, err := h.ApiClient.ListBeers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.renderBeersTemplate(w, beers)
}

func (h *Handler) renderBeersTemplate(w http.ResponseWriter, beers []beverage.Beer) {
	err := h.Templates.ExecuteTemplate(w, beersTemplateFileName, beers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
