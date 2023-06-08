package main

import (
	"errors"
	"fmt"
)

func commandPokedex(state *State, args ...string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments provided.")
	}

	fmt.Println("Pokemons in the Pokedex:")
	for _, pokemon := range state.pokemonsCaught {
		fmt.Printf("+- %v\n", pokemon.Name)
	}
	return nil
}
