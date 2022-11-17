package cmd

import (
	"fmt"
	"godo/utils"
	"strings"

	"github.com/spf13/cobra"
)

// lsCmd represents the lsd command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list out all items in todo.txt",
	Long:  "list out all items in todo.txt",
	Run: func(cmd *cobra.Command, args []string) {
		var todoFile string = utils.TodoDir("todo.txt")
		if len(args) == 0 {
			fmt.Println(utils.ReadFile(todoFile, ""))
		} else {
			fmt.Println(utils.ReadFile(todoFile, strings.Join(args, " ")))
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
