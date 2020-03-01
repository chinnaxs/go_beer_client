package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get all beer from the store",
	Long:  `qewr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func runGet(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(getCmd)
}
