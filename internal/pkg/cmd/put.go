package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/chinnaxs/go_beer_client/internal/pkg/api"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put [name, description, price]",
	Short: "Put more beer in the store",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		runPut(cmd, args)
	},
}

func runPut(cmd *cobra.Command, args []string) {
	var c = &http.Client{}
	price, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println("please specify valid price")
		os.Exit(1)
	}
	beer := &api.Beer{args[0], args[1], price}
	err = api.UpdateBeer(c, beer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("succesfully put beer %s\n", args)
}

func init() {
	RootCmd.AddCommand(putCmd)
}
