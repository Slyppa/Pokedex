package main

import (
	"time"

	"github.com/Slyppa/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pageLocation:  0,
	}
	startRepl(cfg)
}
