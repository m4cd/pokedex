package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCmd struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCmd {
	return map[string]cliCmd{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func repl() {
	fmt.Print("Welcome to pokedex!\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := strings.Fields((scanner.Text()))

		if len(input) == 0 {
			continue
		}

		cmd := strings.ToLower(input[0])

		commands := getCommands()

		command, exists := commands[cmd]

		if exists {
			command.callback()
			continue
		} else {
			fmt.Println("Nonexistent command")
			continue
		}
	}
}
