module github.com/jamistoso/pokedexcli

go 1.23.0

replace github.com/jamistoso/pokedexcli/internal/pokeapi => ./internal/pokeapi
replace github.com/jamistoso/pokedexcli/internal/pokecache => ./internal/pokecache

require github.com/jamistoso/pokedexcli/internal/pokeapi v0.0.0-20241207225842-43e419f5b17e
require github.com/jamistoso/pokedexcli/internal/pokecache