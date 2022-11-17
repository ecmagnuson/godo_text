package cmd

import (
	"godo/utils"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "'do 1' moves line 1 from todo.txt to done.txt. 'do 1 2' moves 1 and 2 from todo.txt to done.txt",
	Long:  "'do 1' moves line 1 from todo.txt to done.txt. 'do 1 2' moves 1 and 2 from todo.txt to done.txt",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, i := range args {
			id, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}
		utils.Do(ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
