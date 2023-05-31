package main

import (
	"fmt"
	"log"
)

func commandMap(state *State, args ...string) error {
	resp, err := state.pokeapiClient.ListLocations(state.next)
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
