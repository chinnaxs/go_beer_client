package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [optional:beername]",
	Short: "get beer from the store",
	Long:  `Lists all beers available if no parameter supplied. Gets details of single beer if name is supplied as parameter.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listBeer(cmd, args)
	},
}

func listBeer(cmd *cobra.Command, args []string) {
	var client = &http.Client{}
	if len(args) > 0 {
		beerName := args[0]
		getOneBeer(client, beerName)
		return
	}
	getAllBeers(client)
}

func getOneBeer(client *http.Client, beerName string) {
	beer, err := api.GetBeer(client, beerName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Succesfully got beer %s\n", beerName)
	fmt.Println(beer)
}

func getAllBeers(client *http.Client) {
	beers, err := api.ListBeers(client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Succesfully got all beers")
	fmt.Println(beers)
}
