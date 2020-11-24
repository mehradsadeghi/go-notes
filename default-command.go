package main

import (
	"os"
)

type DefaultCommand struct {}

func (l DefaultCommand) Handle() {
	showDefaultMessage()
	os.Exit(0)
}