package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jamistoso/pokedexcli/internal/pokeapi"
	"github.com/jamistoso/pokedexcli/internal/pokecache"
)


type locationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type resourceList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type result struct {
	Name		string
	URL			string	
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	index   	int
	offset		int
	next    	string
	previous 	string
	cache	 	pokecache.Cache
}

const locationAreaURL = "https://pokeapi.co/api/v2/location-area/?offset="

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cliCommands()
	cache := pokecache.NewCache(time.Duration(time.Second * 5))
	pokeConfig := config{
		index:    	0,
		offset:	  	20,
		next:     	"",
		previous: 	"",
		cache:		cache,
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
	outStr := "Closing the Pokedex... Goodbye!\n"
	fmt.Println(outStr)
	os.Exit(0)
	return nil
}

func commandMap(conf *config) error {
	val, exists := conf.cache.Get(conf.next)
	if exists {
		location_areas, err := getResults(val)
		if err != nil {
			return fmt.Errorf("result list retrieval failed: %s", err)
		}
		listResultNames(location_areas)
	} else {
		data, err := pokeapi.PokeapiGet(conf.next)
		conf.cache.Add(conf.next, data)
		if err != nil {
			return fmt.Errorf("pokeapi get failed: %s", err)
		}
		location_areas, err := getResults(data)
		if err != nil {
			return fmt.Errorf("location area retrieval failed: %s", err)
		}
		listResultNames(location_areas)
	}
	
	conf.index += conf.offset
	updateConf(conf)
	return nil
}

func commandMapb(conf *config) error {
	if conf.index <= conf.offset {
		outStr := "You're on the first page"
		fmt.Println(outStr)
		return nil
	}
	val, exists := conf.cache.Get(conf.previous)
	if exists {
		location_areas, err := getResults(val)
		if err != nil {
			return fmt.Errorf("result list retrieval failed: %s", err)
		}
		listResultNames(location_areas)
	} else {
		data, err := pokeapi.PokeapiGet(conf.previous)
		conf.cache.Add(conf.previous, data)
		if err != nil {
			return fmt.Errorf("pokeapi get failed: %s", err)
		}
		location_areas, err := getResults(data)
		if err != nil {
			return fmt.Errorf("location area retrieval failed: %s", err)
		}
		listResultNames(location_areas)
	}
	conf.index -= conf.offset
	updateConf(conf)
	return nil
}

func listResultNames(location_areas []result) {
	for _, area := range location_areas {
		fmt.Println(area.Name)
	}
}

func updateConf(conf *config) {
	conf.next = locationAreaURL + strconv.Itoa(conf.index)
	if conf.index >= conf.offset {
		conf.previous = locationAreaURL + strconv.Itoa(conf.index - (conf.offset * 2))
	} else {
		conf.previous = "N/A: 0 index"
	}
}

func getResourceList(data []byte) (resourceList, error) {
	var resList resourceList
	if err := json.Unmarshal([]byte(data), &resList); err != nil {
		return resourceList{}, err
	}
	return resList, nil
}


func getResults(data []byte)  ([]result, error) {

	resList, err := getResourceList(data)
	if err != nil {
		return nil, err
	}

	var location_area_urls []result
	for _, res := range resList.Results {
		item := result{
			Name: res.Name,
			URL: res.URL,
		}
		location_area_urls = append(location_area_urls, item)
	}

	return location_area_urls, nil
}
