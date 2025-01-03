package pokeapi

import (
	"io"
	"net/http"
)

func PokeapiGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}


	return data, nil
}
