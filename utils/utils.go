package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//TodoPath returns the string path (OS agnostic) of the
//todo.txt or done.txt in home/.todo/ dir.
func TodoPath(txtFile string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	} else if txtFile != "todo.txt" && txtFile != "done.txt" {
		panic("incorect file passed into textFile.")
	}
	return filepath.Join(homeDir, ".todo", txtFile)
}

//ReadFile returns a string of the contents of a file given its path.
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		sb.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sb.String()
}

//ReadContexts returns a string array of the contexts in a given file.
func ReadContexts(path string) []string {
	var contexts []string
	return contexts
}
