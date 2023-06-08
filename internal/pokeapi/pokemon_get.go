package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonDetails, error) {
	endpoint := "/pokemon"
	fullURL := baseURL + endpoint + "/" + pokemonName

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return PokemonDetails{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	pokemonDetails := PokemonDetails{}
	err = json.Unmarshal(data, &pokemonDetails)

	if err != nil {
		return PokemonDetails{}, err
	}

	return pokemonDetails, nil
}
