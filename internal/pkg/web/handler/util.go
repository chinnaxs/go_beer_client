package handler

import (
	"net/http"
	"strings"
)

func parseUrl(r *http.Request) string {
	//trimmedUrl := strings.Trim(r.URL.Path, "/")
	tokens := strings.Split(r.URL.Path, "/")
	if len(tokens) > 0 {
		return tokens[len(tokens)-1]
	}
	return ""
}
