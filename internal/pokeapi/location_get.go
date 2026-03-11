package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return Location{}, err
	}

	if data, exists := c.cache.Get(url); exists {
		locationResp := Location{}

		err = json.Unmarshal(data, &locationResp)
		if err != nil {
			return Location{}, nil
		}
		return locationResp, nil
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, nil
	}

	c.cache.Add(url, data)

	return locationResp, nil

}
