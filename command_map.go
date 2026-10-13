package main

import (
	"github.com/Slyppa/Pokedex/internal/Pokeapi"
)

func commandMap(cfg *config) error {
	Pokeapi.Locations(1)
	return nil
}
