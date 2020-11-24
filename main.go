package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	FilePermission = 0666
 	FileName = "notes"
)

type (
	Note struct {
		Title string
		Body  string
	}

	Notes []Note

	Command interface {
		Handle()
	}
)

var Commands = map[string]Command {
	"new": NewCommand{},
	"list": ListCommand{},
}

func main() {
	command := parseCommand(os.Args)
	command.Handle()
}

func parseCommand(args []string) Command {

	if commandIsNotProvidedIn(args) {
		return DefaultCommand{}
	}

	return resolveCommand(fetchCommandName(args))
}

func resolveCommand(command string) Command {

	if _, exists := Commands[command]; exists {
		return Commands[command]
	}

	return DefaultCommand{}
}

func fetchCommandName(args []string) string {
	return args[1]
}

func showDefaultMessage() {
	printHelp()
	os.Exit(0)
}

func commandIsNotProvidedIn(args []string) bool {
	return len(args) < 2
}

func readFile(fileName string) []byte {
	notes, error := ioutil.ReadFile(fileName)
	throwErrorIfNecessary(error)
	return notes
}

func getNotes() Notes {
	notes := make(Notes, 0)
	json.Unmarshal(readFile(FileName), &notes)
	return notes
}

func renderNotes(loadedNotes Notes) {
	for _, note := range loadedNotes {
		fmt.Println("====================")
		fmt.Println("Title: " + note.Title)
		fmt.Println("Body: " + note.Body)
	}
}

func createNote(note Note) {

	collection := append(getNotes(), Note{
		Title: note.Title,
		Body:  note.Body,
	})

	content, error := json.Marshal(collection)

	throwErrorIfNecessary(error)

	writeToFile(content)
}

func writeToFile(output []byte) {
	error := ioutil.WriteFile(FileName, output, FilePermission)
	throwErrorIfNecessary(error)
}

func throwErrorIfNecessary(error error) {
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
}

func getValueFromTerminal(field string) string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(field)
	fmt.Print("-> ")

	value, error := reader.ReadString('\n')

	throwErrorIfNecessary(error)

	// convert CRLF to LF
	value = strings.Replace(value, "\n", "", -1)

	return value
}

func createFileIfNotExists(fileName string) {
	_, error := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, FilePermission)
	throwErrorIfNecessary(error)
}

func printHelp() {
	fmt.Println("Simple Notes!")
	fmt.Println("Try `notes new` to add new note")
	fmt.Println("Try `notes list` to see all notes")
}
