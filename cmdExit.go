package main

import (
	"os"
)

func commandExit(state *State, args ...string) error {
	os.Exit(0)
	return nil
}
