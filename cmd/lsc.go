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
		var todoFile string = utils.TodoPath("todo.txt")
		fmt.Println(utils.ReadContexts(todoFile))
	},
}

func init() {
	rootCmd.AddCommand(lscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
