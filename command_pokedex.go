package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	if len(cfg.caughtPokemon) < 1 {
		return errors.New("you have not caught any pokemons")
	}

	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
