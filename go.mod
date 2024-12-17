module github.com/jamistoso/pokedexcli

go 1.23.0

replace github.com/jamistoso/pokedexcli/internal/pokeapi => ./internal/pokeapi
replace github.com/jamistoso/pokedexcli/internal/pokecache => ./internal/pokecache


require (
	github.com/jamistoso/pokedexcli/internal/pokeapi v0.0.0-20241213211644-0d39163c7644
	github.com/jamistoso/pokedexcli/internal/pokecache v0.0.0-20241213211644-0d39163c7644
)
