package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if data, exists := c.cache.Get(url); exists {
		locationResp := RespShallowLocations{}

		err = json.Unmarshal(data, &locationResp)
		if err != nil {
			return RespShallowLocations{}, nil
		}
		return locationResp, nil
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp := RespShallowLocations{}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, nil
	}

	c.cache.Add(url, data)

	return locationResp, nil

}
