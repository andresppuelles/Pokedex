package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	for k, v := range cfg.commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
