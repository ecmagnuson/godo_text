package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func TodoPath(txtFile string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	} else if txtFile != "todo.txt" && txtFile != "done.txt" {
		panic("incorect file passed into textFile.")
	}
	return filepath.Join(homeDir, ".todo", txtFile)
}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
