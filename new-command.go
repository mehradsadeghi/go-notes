package main

import "os"

type NewCommand struct {}

func (l NewCommand) Handle() {

	createFileIfNotExists(FileName)

	createNote(Note{
		Title: getValueFromTerminal("Title"),
		Body: getValueFromTerminal("Body"),
	})

	os.Exit(0)
}