package main

import (
	"fmt"
	"os"
)

func main() {
	command := parseCommand(os.Args)
	command.Handle()
}

func throwErrorIfNecessary(error error) {
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
}
