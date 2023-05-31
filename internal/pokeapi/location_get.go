package pokeapi

import "fmt"

func (c *Client) GetLocation(name string) (LocationDetails, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + "/" + name
	fmt.Println(fullURL)
	return LocationDetails{}, nil

}
