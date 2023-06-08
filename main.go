package main

import (
	"time"

	"github.com/m4cd/pokedex/internal/pokeapi"
)

type State struct {
	pokeapiClient  pokeapi.Client
	next           *string
	previous       *string
	pokemonsCaught map[string]pokeapi.PokemonDetails
	difficulty     int
}

func main() {
	state := State{
		pokeapiClient:  pokeapi.NewClient(5 * time.Minute),
		pokemonsCaught: map[string]pokeapi.PokemonDetails{},
		difficulty:     40,
	}
	repl(&state)
}
