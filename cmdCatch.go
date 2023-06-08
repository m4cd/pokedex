package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(state *State, args ...string) error {
	if len(args) != 1 {
		return errors.New("Incorrect Pokemon name provided.")
	}
	pokemonName := args[0]

	pokemon, err := state.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	if chance < state.difficulty {
		fmt.Printf("%v was caught!\n", pokemonName)
		state.pokemonsCaught[pokemonName] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}
