package main

import (
	"fmt"
)

func commandHelp(state *State, args ...string) error {
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
