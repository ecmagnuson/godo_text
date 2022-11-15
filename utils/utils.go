package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
//todo.txt or done.txt in $HOME/.todo/ dir.
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
		//ignore empty lines. \r is Windows NT carriage return.
		if strings.TrimSuffix(scanner.Text(), " \r\n") == "" {
			continue
		}
		//if no context given return everything (ls)
		//or
		//if context given return only lines with context (ls @home)
		if (len(context)) == 0 || strings.Contains(scanner.Text(), context) {
			sb.WriteString(scanner.Text() + "\n")
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sb.String()
}

//GetContexts returns a string of all contexts present in a given file.
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

//hasContext checks that text contains a context (@)
func hasContext(text string) bool {
	return strings.Contains(text, "@")
}

//getNumLinesTodo returns the number of lines being used in todo.txt file.
//https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func getNumLinesTodo() (int, error) {

	r := strings.NewReader(ReadFile(TodoPath("todo.txt"), ""))

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

//WriteFile appends text to a file.
func WriteFile(filePath string, text []string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numLines, err := getNumLinesTodo()

	if len(text) > 0 {
		if !hasContext(strings.Join(text, " ")) {
			fmt.Print(text)
			log.Fatal("todo item must have context. It didnt")
		}
		if _, err = f.WriteString("(" + strconv.Itoa(numLines+1) + ") " + strings.Join(text, " ") + "\n"); err != nil {
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
			if _, err = f.WriteString("(" + strconv.Itoa(numLines+1+i) + ") " + text[i]); err != nil {
				panic(err)
			}
		}
	}
}

//getID gets the ID between two parenthesis of a todo item
func getID(s string) int {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			ID, _ := strconv.Atoi(s[i+1 : j])
			return ID
		}
	}
	panic("getID can't find an ID in a todo item")
}

//Do moves the id (line) in todo.txt to done.txt and rewrites the line in todo.txt to ""
func Do(ids []int) {
	var stringOfTodos string = ReadFile(TodoPath("todo.txt"), "")
	var todos []string
	//convert todos to an array of strings per line?
	scanner := bufio.NewScanner(strings.NewReader(stringOfTodos))
	for scanner.Scan() {
		todos = append(todos, scanner.Text())
	}

	//I have a []string - todos - of all of the things todo
	//The parameter ids []int is all of the items I want to do

	//I need to get a list of all of the things in ids that are in todos

	var done []string

	i := 0
	for i < len(todos) {
		if slices.Contains(ids, getID(todos[i])) {
			fmt.Println(todos[i])
		}
		i++
	}

	/* 	for i := 1; i < len(todos); i++ {
		if slices.Contains(ids, i) && todos[i] != "" {
			todos[i] = strings.TrimSuffix(todos[i], " \r\n")
			done = append(done, todos[i-1])
		}
	} */

	fmt.Println()
	fmt.Println("Todos are:")
	fmt.Println(todos)
	fmt.Println()
	fmt.Println("Done is:")
	fmt.Println(done)
}
