package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (LocationDetails, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + "/" + name

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationDetails{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	locationDetails := LocationDetails{}
	err = json.Unmarshal(data, &locationDetails)

	if err != nil {
		return LocationDetails{}, err
	}

	return locationDetails, nil

}
