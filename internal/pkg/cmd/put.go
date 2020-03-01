package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")
	},
}

func runPut(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(putCmd)
}
