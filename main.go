package main

import (
	"time"

	"github.com/Slyppa/Pokedex/internal/Pokeapi"
)

func main() {
	pokeClient := Pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
