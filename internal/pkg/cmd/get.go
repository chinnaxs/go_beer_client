package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [opt:beername",
	Short: "get beer from the store",
	Long:  `Lists all beers available if no parameter supplied. Gets details of single beer if name is supplied as parameter.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGet(cmd, args)
	},
}

func runGet(cmd *cobra.Command, args []string) {
	var c = &http.Client{}
	if len(args) > 0 {
		beerName := args[0]
		beer, err := api.GetBeer(c, beerName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Succesfully got beer %s\n", beerName)
		fmt.Println(beer)
		return
	}

	beers, err := api.ListBeers(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Succesfully got all beers")
	fmt.Println(beers)
}

func init() {
	RootCmd.AddCommand(getCmd)
}
