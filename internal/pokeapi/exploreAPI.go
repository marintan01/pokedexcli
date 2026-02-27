package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/marintan01/pokedexcli/internal/pokecache"
	"io"
	"log"
	"net/http"
)

func GetPokemonInLocation(pokecache *pokecache.Cache, location string) (PokemonSingleLocation, error) {

	full_url := baseUrl + "/location-area/" + location

	if val, ok := pokecache.Get(full_url); ok {
		locationsResp := PokemonSingleLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokemonSingleLocation{}, err
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
		return PokemonSingleLocation{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokemonSingleLocation{}, err
	}

	singleLocation := PokemonSingleLocation{}
	err = json.Unmarshal(body, &singleLocation)
	if err != nil {
		return PokemonSingleLocation{}, err
	}

	pokecache.Add(full_url, body)

	return singleLocation, nil

}
