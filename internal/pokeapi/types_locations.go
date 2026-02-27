package pokeapi

type PokemonLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonSingleLocation struct {
	Id                int     `json:"id"`
	Name              *string `json:"name"`
	GameIndex         int     `json:"game_index"`
	PokemonEncaunters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
