package main

import (
	"errors"
	"fmt"
)

func commandExplore(state *State, args ...string) error {
	if len(args) != 1 {
		return errors.New("Location value not provided.")
	}
	locationName := args[0]

	l, err := state.pokeapiClient.GetLocation(locationName)

	if err != nil {
		return err
	}

	for _, pokemon := range l.PokemonEncounters {
		fmt.Printf("+- %v\n", pokemon)
	}
	return nil
}
