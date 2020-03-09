package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/chinnaxs/go_beer_client/internal/pkg/beverage"
)

var updateTemplateFileName = "update.html"

var updateHandler *Handler

func MakeUpdateHandler(templatePath string, apiClient *api.ApiClient) http.HandlerFunc {

	beersTemplatePath := filepath.Join(templatePath, beersTemplateFileName)
	updateHandler = &Handler{
		Templates: template.Must(template.ParseFiles(
			beersTemplatePath,
		)),
		ApiClient: apiClient,
	}
	return updateHandler.updateBeerHandler
}

func (h *Handler) updateBeerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUpdateBeerHandler(w, r)
	case http.MethodPost:
		h.postUpdateBeerHandler(w, r)
	case http.MethodDelete:
		h.deleteUpdateBeerHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) getUpdateBeerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	beerName := parseUrl(r)
	if beerName == "" {
		h.newBeerHandler(w, r)
	} else {
		h.existingBeerHandler(w, r, beerName)
	}
}

func (h *Handler) postUpdateBeerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	description := r.FormValue("Description")
	priceString := r.FormValue("Price")
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Body: %s, %f, %s", name, price, description)
	beer := &beverage.Beer{
		Name:        name,
		Description: description,
		Price:       price,
	}
	err = h.ApiClient.UpdateBeer(beer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/beer/"+name, http.StatusFound)
}

func (h *Handler) deleteUpdateBeerHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) newBeerHandler(w http.ResponseWriter, r *http.Request) {
	beer := &beverage.Beer{
		Name:        "",
		Description: "",
		Price:       5,
	}
	h.renderUpdateTemplate(w, beer)
}

func (h *Handler) existingBeerHandler(w http.ResponseWriter, r *http.Request, beerName string) {
	beer, err := h.ApiClient.GetBeer(beerName)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	h.renderUpdateTemplate(w, beer)
}

func (h *Handler) renderUpdateTemplate(w http.ResponseWriter, beer *beverage.Beer) {
	err := h.Templates.ExecuteTemplate(w, beersTemplateFileName, beer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
