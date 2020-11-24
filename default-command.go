package main

import (
	"fmt"
	"os"
)

type DefaultCommand struct {}

func (l DefaultCommand) Handle() {
	showDefaultMessage()
	os.Exit(0)
}

func showDefaultMessage() {
	fmt.Println("Simple Notes!")
	fmt.Println("Try `notes new` to add new note")
	fmt.Println("Try `notes list` to see all notes")
	os.Exit(0)
}