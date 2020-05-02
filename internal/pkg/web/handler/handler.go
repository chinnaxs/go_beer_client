package handler

import (
	"html/template"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
)

type Handler struct {
	Templates    *template.Template
	ApiClient    *api.ApiClient
}
