package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
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

//ReadFile returns a string of all of the contents of a file given its path.
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	i := -1 //TODO: fix this
	for scanner.Scan() {
		i++
		if scanner.Text() == "" {
			continue
		}
		sb.WriteString("(" + strconv.Itoa(i) + ") " + scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sb.String()
}

//ReadContext reads only specific contexts from a file.
func ReadContext(path string, context string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	i := -1 //TODO: fix this
	for scanner.Scan() {
		i++
		if scanner.Text() == "" {
			continue
		}
		if strings.Contains(scanner.Text(), context) {
			sb.WriteString("(" + strconv.Itoa(i) + ") " + scanner.Text() + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sb.String()
}

//GetContexts returns a string array of the contexts present in a given file.
func GetContexts(path string) string {
	var contexts []string
	var todos string = ReadFile(TodoPath("todo.txt"))
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

func hasContext(text string) bool {
	return strings.Contains(text, "@")
}

//WriteFile writes text to a file.
func WriteFile(filePath string, text []string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if len(text) > 0 {
		if !hasContext(strings.Join(text, " ")) {
			log.Fatal("todo item must have context.")
		}
		if _, err = f.WriteString(strings.Join(text, " ") + "\n"); err != nil {
			panic(err)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			next, _ := reader.ReadString('\n')
			if next == "\n" {
				break
			}
			if !hasContext(next) {
				fmt.Println("todo item must have context.")
				continue
			}
			text = append(text, next)
		}
		for i := 0; i < len(text); i++ {
			if _, err = f.WriteString(text[i]); err != nil {
				panic(err)
			}
		}
	}
}

//Do moves the id (line) in todo.txt to done.txt
func Do(ids []int) {
	var todos string = ReadFile(TodoPath("todo.txt"))
	var todoList []string
	//convert todos to an array of strings per line?
	scanner := bufio.NewScanner(strings.NewReader(todos))
	for scanner.Scan() {
		todoList = append(todoList, scanner.Text())
	}
	var done []int
	for i := 0; i < len(todos); i++ {
		if slices.Contains(ids, i) {

		}
	}
}
