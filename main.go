package main

import (
	"github.com/andresppuelles/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	cfg := config{
		commands: map[string]cliCommand{},
		cache: pokecache.NewCache(50*time.Second),
	}
	startRepl(&cfg)
}
