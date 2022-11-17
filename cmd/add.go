package cmd

import (
	"godo/utils"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	Short: `everything after the 'add' command will be added to todo.txt.
			'add' with no commands after will allow you to add multiple items at once
			When you are done, press 'enter' to add all of the lines.`,
	Long: `everything after the 'add' command will be added to todo.txt.
			'add' with no commands after will allow you to add multiple items at once
			When you are done, press 'enter' to add all of the lines.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.WriteFile(utils.TodoDir("todo.txt"), args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
