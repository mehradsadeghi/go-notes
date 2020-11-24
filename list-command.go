package main

import (
	"os"
)

type ListCommand struct {}

func (l ListCommand) Handle() {
	createFileIfNotExists(FileName)
	renderNotes(getNotes())
	os.Exit(0)
}