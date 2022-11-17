package cmd

import (
	"fmt"
	"godo/utils"

	"github.com/spf13/cobra"
)

// lscCmd represents the lsc command
var lscCmd = &cobra.Command{
	Use:   "lsc",
	Short: "list contexts",
	Long:  "list contexts",
	Run: func(cmd *cobra.Command, args []string) {
		var todoFile string = utils.TodoDir("todo.txt")
		fmt.Println(utils.GetContexts(todoFile))
	},
}

func init() {
	rootCmd.AddCommand(lscCmd)
}
