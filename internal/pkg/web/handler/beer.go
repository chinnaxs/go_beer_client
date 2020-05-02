package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/chinnaxs/go_beer_client/internal/pkg/beverage"
)

var beerTemplateFileName = "beer.html"
var beerHandler *Handler

func BeerHandler(templatePath string, apiClient *api.ApiClient) http.HandlerFunc {
	beersTemplatePath := filepath.Join(templatePath, beerTemplateFileName)
	beerHandler := Handler{
		Templates: template.Must(template.ParseFiles(
			beersTemplatePath,
		)),
		ApiClient: apiClient,
	}
	return beerHandler.singleBeerHandler
}

func (h *Handler) singleBeerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	if r.Method == http.MethodGet {
		beerHandler.getBeerHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (h *Handler) getBeerHandler(w http.ResponseWriter, r *http.Request) {
	beerName := parseUrl(r)
	if beerName == "" {
		http.NotFound(w, r)
		return
	}
	beer, err := h.ApiClient.GetBeer(beerName)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	h.renderBeerTemplate(w, beer)
}

func (h *Handler) renderBeerTemplate(w http.ResponseWriter, beer *beverage.Beer) {
	err := h.Templates.ExecuteTemplate(w, beerTemplateFileName, beer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
