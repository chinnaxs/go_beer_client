package web

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/chinnaxs/go_beer_client/internal/pkg/web/handler"
)

var rootPath, _ = os.Executable()
var templatePath = filepath.Join(rootPath, "..", "..", "..", "web", "template")

type Server struct {
	Addr      *net.TCPAddr
	ApiClient *api.ApiClient
}

func NewServer(Addr *net.TCPAddr) *Server {
	return &Server{
		Addr:      Addr,
		ApiClient: api.NewDefaultApiClient(),
	}
}

func (s *Server) Start() {
	log.Printf("Loaded templates from: %s\n", templatePath)

	var beersHandler = handler.MakeBeersHandler(templatePath, s.ApiClient)
	var beerHandler = handler.MakeBeerHandler(templatePath, s.ApiClient)
	var updateHandler = handler.MakeUpdateHandler(templatePath, s.ApiClient)
	http.HandleFunc("/", beersHandler)
	http.HandleFunc("/beer/", beerHandler)
	http.HandleFunc("/update/", updateHandler)

	//http.HandleFunc("/edit/", s.editHandler)
	log.Printf("Start listening on %s\n", s.Addr.String())
	log.Fatal(http.ListenAndServe(s.Addr.String(), nil))
}

// func (s *Server) editHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("%s: %s", r.Method, r.URL.Path)
// 	switch r.Method {
// 	case http.MethodGet:
// 		s.handleEditGetBeer(w, r)
// 	case http.MethodPost:
// 		s.handleEditPostBeer(w, r)
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }

// func (s *Server) handleEditGetBeer(w http.ResponseWriter, r *http.Request) {
// 	beerName := parseUrl(r)
// 	beers, err := s.ApiClient.GetBeer(beerName)
// 	if err != nil {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	renderEditBeerTemplate(w, beers)
// }

// func (s *Server) handleEditPostBeer(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("editpost called")
// }

// func (s *Server) handlePutBeer(w http.ResponseWriter, r *http.Request) {

// }

// func (s *Server) handleDeleteBeer(w http.ResponseWriter, r *http.Request) {
// }

// func renderEditBeerTemplate(w http.ResponseWriter, beer *beverage.Beer) {
// 	err := templates.ExecuteTemplate(w, editTemplateFileName, beer)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
