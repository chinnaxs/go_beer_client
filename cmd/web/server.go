package main

import (
	"log"
	"net"

	"github.com/chinnaxs/go_beer_client/internal/pkg/web"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatal("was not able to parse tcp Addr")
	}
	s := web.NewServer(tcpAddr)
	s.Start()
}
