package main

type Command interface {
	Handle()
}

var Commands = map[string]Command {
	"new": NewCommand{},
	"list": ListCommand{},
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

func commandIsNotProvidedIn(args []string) bool {
	return len(args) < 2
}