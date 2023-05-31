package main

import (
	"time"

	"github.com/m4cd/pokedex/internal/pokeapi"
)

type State struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func main() {
	state := State{
		pokeapiClient: pokeapi.NewClient(5 * time.Minute),
	}
	repl(&state)
}
