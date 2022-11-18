package cmd

import (
	"godo/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//setup creates .todo dir in user home and puts todo.txt and done.txt inside of it
func setup() {
	todoPath := utils.TodoDir("todo.txt")
	donePath := utils.TodoDir("done.txt")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	newpath := filepath.Join(homeDir, ".todo")

	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	td, err := os.Create(todoPath)
	if err != nil {
		panic(err)
	}

	d, err := os.Create(donePath)
	if err != nil {
		panic(err)
	}

	defer td.Close()
	defer d.Close()
}

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup the .todo dir in user $HOME and populate a todo.txt and done.txt file",
	Long:  "setup the .todo dir in user $HOME and populate a todo.txt and done.txt file",
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
