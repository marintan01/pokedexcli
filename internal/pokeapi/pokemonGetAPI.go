package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/marintan01/pokedexcli/internal/pokecache"
	"io"
	"log"
	"net/http"
)

func GetPokemon(pokecache *pokecache.Cache, pokemonName string) (Pokemon, error) {

	full_url := baseUrl + "/pokemon/" + pokemonName

	if val, ok := pokecache.Get(full_url); ok {
		locationsResp := Pokemon{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Pokemon{}, err
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
		return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	pokecache.Add(full_url, body)

	return pokemon, nil

}
