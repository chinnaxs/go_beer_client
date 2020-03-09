package cmd

import (
	"fmt"
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
	a := api.NewDefaultApiClient()
	if len(args) > 0 {
		beerName := args[0]
		getOneBeer(a, beerName)
		return
	}
	getAllBeers(a)
}

func getOneBeer(a *api.ApiClient, beerName string) {
	beer, err := a.GetBeer(beerName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Succesfully got beer %s\n", beerName)
	fmt.Println(beer)
}

func getAllBeers(a *api.ApiClient) {
	beers, err := a.ListBeers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Succesfully got all beers")
	fmt.Println(beers)
}
