package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetLocations(pageUrl *string) (PokemonLocations, error) {

	full_url := baseUrl + "/location-area"
	if pageUrl != nil {
		full_url = *pageUrl
	}
	res, err := http.Get(full_url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonLocations{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokemonLocations{}, err
	}

	location := PokemonLocations{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return PokemonLocations{}, err
	}

	return location, nil

}
