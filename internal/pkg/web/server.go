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
	var deleteHandler = handler.MakeDeleteBeerHandler(templatePath, s.ApiClient)
	http.HandleFunc("/", beersHandler)
	http.HandleFunc("/beer/", beerHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/delete/", deleteHandler)

	log.Printf("Start listening on %s\n", s.Addr.String())
	log.Fatal(http.ListenAndServe(s.Addr.String(), nil))
}
