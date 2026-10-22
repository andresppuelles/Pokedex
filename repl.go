package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/andresppuelles/pokedexcli/internal/pokecache"
)

func startRepl(cfg *config) {
	buildCommands(cfg)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		value, ok := (*cfg).commands[commandName]
		if ok {
			err := value.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleaned := []string{}
	if len(text) < 1 {
		return []string{}
	}
	stripped := strings.Fields(text)
	for _, word := range stripped {
		cleaned = append(cleaned, strings.ToLower(word))
	}
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextLocationURL *string
    prevLocationURL *string
    cache pokecache.Cache
	commands map[string]cliCommand
}

func buildCommands(cfg *config) {
	supportedCommand := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map":{
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world, each subsequent call displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb":{
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world, each subsequent call displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
	cfg.commands = supportedCommand
}
