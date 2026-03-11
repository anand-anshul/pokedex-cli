package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	if data, exists := c.cache.Get(url); exists {
		pokemonResp := Pokemon{}

		err = json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, nil
		}
		return pokemonResp, nil
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(url, data)

	return pokemonResp, nil

}
