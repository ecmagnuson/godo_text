package cmd

import (
	"godo/utils"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup the .todo dir in user $HOME and populate a todo.txt and done.txt file",
	Long:  "setup the .todo dir in user $HOME and populate a todo.txt and done.txt file",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
