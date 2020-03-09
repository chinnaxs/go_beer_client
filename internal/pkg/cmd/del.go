package cmd

import (
	"fmt"
	"os"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use:   "del [name]",
	Short: "delete beer from the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		removeBeer(cmd, args)
	},
}

func removeBeer(cmd *cobra.Command, args []string) {
	a := api.NewDefaultApiClient()
	beerName := args[0]
	err := a.DeleteBeer(beerName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("succesfully deleted beer %s\n", beerName)
}
