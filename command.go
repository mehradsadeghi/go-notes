package main

const MinArgLength = 2

type Command interface {
	Handle()
}

var Commands = map[string]Command {
	"new": NewCommand{},
	"list": ListCommand{},
	"default": DefaultCommand{},
}

func parseCommand(args []string) Command {

	if commandIsNotProvidedIn(args) {
		return Commands["default"]
	}

	return resolveCommand(fetchCommandName(args))
}

func resolveCommand(command string) Command {

	if _, exists := Commands[command]; exists {
		return Commands[command]
	}

	return Commands["default"]
}

func fetchCommandName(args []string) string {
	return args[1]
}

func commandIsNotProvidedIn(args []string) bool {
	return len(args) < MinArgLength
}