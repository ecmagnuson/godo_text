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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//lsCmd.Flags().BoolP("toggle", "@", false, "Help message for toggle")
	//lsCmd.MarkFlagRequired("toggle")
}
