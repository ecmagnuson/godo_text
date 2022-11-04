/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
		var todoFile string = utils.TodoPath("done.txt")
		fmt.Println(utils.ReadFile(todoFile))
	},
}

func init() {
	rootCmd.AddCommand(lsdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
