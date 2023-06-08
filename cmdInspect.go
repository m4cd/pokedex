package main

import (
	"errors"
	"fmt"
)

func commandInspect(state *State, args ...string) error {
	if len(args) != 1 {
		return errors.New("Incorrect Pokemon name provided.")
	}
	pokemonName := args[0]

	pokemon, ok := state.pokemonsCaught[pokemonName]

	if !ok {
		return fmt.Errorf("You haven't caught %v yet...", pokemonName)
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("+- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("+- %v\n", t.Type.Name)
	}
	return nil
}
