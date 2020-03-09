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

func MakeBeersHandler(templatePath string, apiClient *api.ApiClient) http.HandlerFunc {

	beersTemplatePath := filepath.Join(templatePath, beersTemplateFileName)
	beersHandler = &Handler{
		Templates: template.Must(template.ParseFiles(
			beersTemplatePath,
		)),
		ApiClient: apiClient,
	}
	return beersHandler.listBeerHandler
}

func (h *Handler) listBeerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		beersHandler.handleGetAllBeers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) handleGetAllBeers(w http.ResponseWriter, r *http.Request) error {
	beers, err := h.ApiClient.ListBeers()
	if err != nil {
		return err
	}
	h.renderBeersTemplate(w, beers)
	return nil
}

func (h *Handler) renderBeersTemplate(w http.ResponseWriter, beers []beverage.Beer) {
	err := h.Templates.ExecuteTemplate(w, beersTemplateFileName, beers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
