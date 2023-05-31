package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(page *string) (LocationsResponse, error) {
	endpoint := "/location-area"

	fullURL := baseURL + endpoint
	if page != nil {
		fullURL = *page
	}

	//c.cache.get()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationsResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	locationsResponse := LocationsResponse{}
	err = json.Unmarshal(data, &locationsResponse)

	if err != nil {
		return LocationsResponse{}, err
	}
	return locationsResponse, err
}
