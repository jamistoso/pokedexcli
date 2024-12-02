package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	id         int
	name       string
	game_index int
	location   NamedAPIResource
}

type NamedAPIResource struct {
	name string
	url  string
}

func getLocationAreas(url string) ([]LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var location_areas []LocationArea
	if err = json.Unmarshal(data, &location_areas); err != nil {
		return nil, err
	}

	return location_areas, nil
}
