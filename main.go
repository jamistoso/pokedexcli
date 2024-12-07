package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jamistoso/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	index    int
	offset	 int
	next     string
	previous string
}

const locationAreaURL = "https://pokeapi.co/api/v2/location-area/?offset="

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cliCommands()
	pokeConfig := config{
		index:    0,
		offset:	  20,
		next:     "",
		previous: "",
	}
	updateConf(&pokeConfig)
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
			err := function.callback(&pokeConfig)
			if err != nil {
				fmt.Println(err)
				continue
			}
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
	var location_areas []pokeapi.LocationAreaURL
	location_areas, err := pokeapi.GetLocationAreaURLs(conf.next)
	if err != nil {
		return fmt.Errorf("location error retrieval failed: %s", err)
	}
	listAreaURLNames(location_areas)
	conf.index += conf.offset
	updateConf(conf)
	return nil
}

func commandMapb(conf *config) error {
	if conf.index == 0 {
		outStr := "No previous location areas found"
		fmt.Println(outStr)
		return nil
	}
	var location_areas []pokeapi.LocationAreaURL
	location_areas, err := pokeapi.GetLocationAreaURLs(conf.previous)
	if err != nil {
		return fmt.Errorf("location error retrieval failed: %s", err)
	}
	listAreaURLNames(location_areas)
	conf.index -= conf.offset
	updateConf(conf)
	return nil
}

func listAreaURLNames(location_areas []pokeapi.LocationArea) {
	for _, area := range location_areas {
		fmt.Println(area.Name)
	}
}

func updateConf(conf *config) {
	conf.next = locationAreaURL + strconv.Itoa(conf.index)
	if conf.index != 0 {
		conf.previous = locationAreaURL + strconv.Itoa(conf.index - conf.offset)
	} else {
		conf.previous = "N/A: 0 index"
	}
}
