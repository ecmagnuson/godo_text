package cmd

import (
	"bufio"
	"fmt"
	"godo/utils"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

//GetContexts returns a string of all contexts present in a given file.
func GetContexts(path string) string {
	var contexts []string
	var todos string = utils.ReadFile(utils.TodoDir("todo.txt"), "")
	scanner := bufio.NewScanner(strings.NewReader(todos))
	for scanner.Scan() {
		var line string = scanner.Text()
		line = line[strings.IndexByte(line, '@'):]
		if !slices.Contains(contexts, line) {
			contexts = append(contexts, line)
		}
	}
	return strings.Join(contexts, " ")
}

// lscCmd represents the lsc command
var lscCmd = &cobra.Command{
	Use:   "lsc",
	Short: "list contexts",
	Long:  "list contexts",
	Run: func(cmd *cobra.Command, args []string) {
		var todoFile string = utils.TodoDir("todo.txt")
		fmt.Println(GetContexts(todoFile))
	},
}

func init() {
	rootCmd.AddCommand(lscCmd)
}
