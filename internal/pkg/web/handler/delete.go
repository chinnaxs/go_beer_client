package handler

import (
	"log"
	"net/http"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
)

var deleteHandler *Handler

func DeleteBeerHandler(templatePath string, apiClient *api.ApiClient) http.HandlerFunc {
	deleteHandler := Handler{
		ApiClient: apiClient,
	}
	return deleteHandler.deleteBeerHandler
}

func (h *Handler) deleteBeerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	if r.Method == http.MethodPost {
		h.postDeleteBeerHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
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
