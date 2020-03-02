package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [name]",
	Short: "delete beer from the store",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runDel(cmd, args)
	},
}

func runDel(cmd *cobra.Command, args []string) {
	var c = &http.Client{}
	beerName := args[0]
	err := api.DeleteBeer(c, beerName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("succesfully deleted beer %s\n", beerName)
}

func init() {
	RootCmd.AddCommand(delCmd)
}