package main

import (
	"github.com/Slyppa/Pokedex/internal/Pokeapi"
)

func main() {
	/*cfg := &config{
		commands: getCommands(),
	}
	startRepl(cfg)*/
	Pokeapi.Locations(1)
}
