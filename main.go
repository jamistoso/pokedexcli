package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       any    `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  any    `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      any    `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
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
	callback    func(*config, string) error
}

type config struct {
	index   	int
	offset		int
	next    	string
	previous 	string
	cache	 	pokecache.Cache
}

const locationAreaListURL = "https://pokeapi.co/api/v2/location-area/?offset="
const locationAreaURL = 	"https://pokeapi.co/api/v2/location-area/"
const pokemonURL =			"https://pokeapi.co/api/v2/pokemon/"

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
			args := strings.Split(command, " ")
			function, ok := commands[args[0]]
			if !ok {
				fmt.Println("unknown command: " + command)
				continue
			}
			var err error
			if len(args) == 1 {
				err = function.callback(&pokeConfig, "")
			} else {
				err = function.callback(&pokeConfig, args[1])
			}
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
		"explore": {
			name:        "explore",
			description: "Retrieve a list of pokemon within a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
	}
}

func commandHelp(conf *config, arg1 string) error {
	outStr := "Welcome to the Pokedex!\nUsage:\n\n"
	commands := cliCommands()
	for command := range commands {
		outStr += commands[command].name + ": " + commands[command].description + "\n"
	}
	fmt.Println(outStr)
	return nil
}

func commandExit(conf *config, arg1 string) error {
	outStr := "Closing the Pokedex... Goodbye!\n"
	fmt.Println(outStr)
	os.Exit(0)
	return nil
}

func commandMap(conf *config, arg1 string) error {
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

func commandMapb(conf *config, arg1 string) error {
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

func commandExplore(conf *config, arg1 string) error {
	url := locationAreaURL + arg1
	val, exists := conf.cache.Get(url)
	if exists {
		location_area, err := getLocationArea(val)
		if err != nil {
			return fmt.Errorf("result list retrieval failed: %s", err)
		}
		listPokemonInLocationArea(location_area)
	} else {
		data, err := pokeapi.PokeapiGet(url)
		conf.cache.Add(conf.next, data)
		if err != nil {
			return fmt.Errorf("pokeapi get failed: %s", err)
		}
		location_area, err := getLocationArea(data)
		if err != nil {
			return fmt.Errorf("location area retrieval failed: %s", err)
		}
		listPokemonInLocationArea(location_area)
	}
	
	conf.index += conf.offset
	updateConf(conf)
	return nil
}

func commandCatch(conf *config, arg1 string) error {
	url := locationAreaURL + arg1
	val, exists := conf.cache.Get(url)
	if exists {
		location_area, err := getLocationArea(val)
		if err != nil {
			return fmt.Errorf("result list retrieval failed: %s", err)
		}
		listPokemonInLocationArea(location_area)
	} else {
		data, err := pokeapi.PokeapiGet(url)
		conf.cache.Add(conf.next, data)
		if err != nil {
			return fmt.Errorf("pokeapi get failed: %s", err)
		}
		location_area, err := getLocationArea(data)
		if err != nil {
			return fmt.Errorf("location area retrieval failed: %s", err)
		}
		listPokemonInLocationArea(location_area)
	}
	
	conf.index += conf.offset
	updateConf(conf)
	return nil
}

func listResultNames(location_areas []result) {
	for _, area := range location_areas {
		fmt.Println(area.Name)
	}
}

func listPokemonInLocationArea(location_area locationArea) {
	for _, encounter := range location_area.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}

func updateConf(conf *config) {
	conf.next = locationAreaListURL + strconv.Itoa(conf.index)
	if conf.index >= conf.offset {
		conf.previous = locationAreaListURL + strconv.Itoa(conf.index - (conf.offset * 2))
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

func getLocationArea(data []byte) (locationArea, error) {
	var locArea locationArea
	if err := json.Unmarshal([]byte(data), &locArea); err != nil {
		return locationArea{}, err
	}
	return locArea, nil
}