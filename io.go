package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	FilePermission = 0666
	FileName = "notes"
)

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

func writeToFile(output []byte) {
	error := ioutil.WriteFile(FileName, output, FilePermission)
	throwErrorIfNecessary(error)
}

func readFile(fileName string) []byte {
	notes, error := ioutil.ReadFile(fileName)
	throwErrorIfNecessary(error)
	return notes
}