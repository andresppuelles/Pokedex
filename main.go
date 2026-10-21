package main

func main() {
	cfg := config{
		commands: map[string]cliCommand{},
	}
	startRepl(&cfg)
}
