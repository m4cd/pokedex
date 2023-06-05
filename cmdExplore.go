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
	fmt.Printf("Pokemons found in %v:\n", locationName)
	for _, pokemon := range l.PokemonEncounters {
		fmt.Printf("+- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}
