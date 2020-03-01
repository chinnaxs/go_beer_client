package main

import (
	"log"

	"github.com/chinnaxs/go_beer_client/internal/pkg/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
