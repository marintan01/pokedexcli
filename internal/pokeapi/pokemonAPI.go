package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/marintan01/pokedexcli/internal/pokecache"
	"io"
	"log"
	"net/http"
)

func GetLocations(pokecache *pokecache.Cache, pageUrl *string) (PokemonLocations, error) {

	full_url := baseUrl + "/location-area"
	if pageUrl != nil {
		full_url = *pageUrl
	}

	if val, ok := pokecache.Get(full_url); ok {
		locationsResp := PokemonLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokemonLocations{}, err
		}

		return locationsResp, nil
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
