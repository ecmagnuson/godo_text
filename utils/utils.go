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

//TODO
//When adding to todo.txt need to only append
//Same with done.txt
//when using Do want to fully rewrite everything

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

//ReadFile returns a string of all of the contents of a file given its path if no context is given.
//If there is a context given it will return only the lines with that context.
func ReadFile(path string, context string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		if strings.TrimSuffix(scanner.Text(), " \r\n") == "" {
			continue
		}
		if (len(context)) == 0 || strings.Contains(scanner.Text(), context) {
			sb.WriteString("(" + strconv.Itoa(i) + ") " + scanner.Text() + "\n")
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sb.String()
}

//GetContexts returns a string of the contexts present in a given file.
func GetContexts(path string) string {
	var contexts []string
	var todos string = ReadFile(TodoPath("todo.txt"), "")
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

//WriteFile appends text to a file.
func WriteFile(filePath string, text []string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if len(text) > 0 {
		if !hasContext(strings.Join(text, " ")) {
			fmt.Print(text)
			log.Fatal("todo item must have context. It didnt")
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
	var todoString string = ReadFile(TodoPath("todo.txt"), "")
	var todos []string
	//convert todos to an array of strings per line?
	scanner := bufio.NewScanner(strings.NewReader(todoString))
	for scanner.Scan() {
		todos = append(todos, scanner.Text())
	}

	//I have a []string - todos - of all of the things todo
	//The parameter ids []int is all of the items I want to do

	//I need to get a list of all of the things in ids that are in todos

	var done []string

	for i := 0; i < len(todos); i++ {
		if slices.Contains(ids, i) && todos[i] != "" {
			todos[i] = strings.TrimSuffix(todos[i], " \r\n")
			done = append(done, todos[i-1]) //This is a BUG workaround.
		}
	}
	fmt.Println(done)
}
