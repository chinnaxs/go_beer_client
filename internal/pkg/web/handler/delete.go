package handler

import (
	"log"
	"net/http"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
)

var deleteHandler *Handler

func MakeDeleteBeerHandler(templatePath string, apiClient *api.ApiClient) http.HandlerFunc {
	deleteHandler := Handler{
		Templates: nil,
		ApiClient: apiClient,
	}
	return deleteHandler.deleteBeerHandler
}

func (h *Handler) deleteBeerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		h.postDeleteBeerHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) postDeleteBeerHandler(w http.ResponseWriter, r *http.Request) {
	beerName := parseUrl(r)
	err := h.ApiClient.DeleteBeer(beerName)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/beers", http.StatusFound)
}
