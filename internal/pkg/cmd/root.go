package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "Beer store CLI",
	Short: "Beer store CLI for go_beer API",
	Long:  "",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//  Run: func(cmd *cobra.Command, args []string) { },
}
