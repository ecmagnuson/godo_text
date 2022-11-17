package cmd

import (
	"fmt"
	"godo/utils"

	"github.com/spf13/cobra"
)

// lsdCmd represents the lsd command
var lsdCmd = &cobra.Command{
	Use:   "lsd",
	Short: "list items in done.txt",
	Long:  "list items in done.txt",
	Run: func(cmd *cobra.Command, args []string) {
		var todoFile string = utils.TodoDir("done.txt")
		fmt.Println(utils.ReadFile(todoFile, ""))
	},
}

func init() {
	rootCmd.AddCommand(lsdCmd)
}
