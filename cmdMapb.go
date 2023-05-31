package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMapb(state *State, args ...string) error {

	if state.previous == nil {
		return errors.New("That's first page. Cannot go back.")
	}

	resp, err := state.pokeapiClient.ListLocations(state.previous)
	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("Locations:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	state.next = resp.Next
	state.previous = resp.Previous

	return nil
}
