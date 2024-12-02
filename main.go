package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	index    int
	next     string
	previous string
}

const locationAreaURL = "https://pokeapi.co/api/v2/location-area/?offset="

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cliCommands()
	pokeConfig := config{
		index:    0,
		next:     "",
		previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if scanner.Err() == nil {
				continue
			}

			// TODO: Implement command error handling
		} else {
			command := scanner.Text()
			function, ok := commands[command]
			if !ok {
				fmt.Println("unknown command")
				continue
			}
			function.callback(&pokeConfig)
		}

	}
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Retrieve the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Retrieve the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func commandHelp(conf *config) error {
	outStr := "Welcome to the Pokedex!\nUsage:\n\n"
	commands := cliCommands()
	for command := range commands {
		outStr += commands[command].name + ": " + commands[command].description + "\n"
	}
	fmt.Println(outStr)
	return nil
}

func commandExit(conf *config) error {
	os.Exit(0)
	return nil
}

func commandMap(conf *config) error {
	return nil
}

func commandMapb(conf *config) error {
	if conf.index == 0 {
		outStr := "No previous location areas found"
		fmt.Println(outStr)
		return nil
	}
	conf.index -= 20
	return nil
}
