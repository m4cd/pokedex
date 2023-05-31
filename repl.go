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
	callback    func(*State, ...string) error
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
		"map": {
			name:        "map",
			description: "Prints next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explores provided location",
			callback:    commandExplore,
		},
	}
}

func repl(state *State) {
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
		args := []string{}

		if len(input) > 1 {
			args = input[1:]
		}

		commands := getCommands()
		command, exists := commands[cmd]

		if exists {
			err := command.callback(state, args...)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			continue
		} else {
			fmt.Println("Non-existent command")
			continue
		}
	}
}
