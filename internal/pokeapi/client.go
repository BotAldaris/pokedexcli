package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BotAldaris/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, clearCache time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(clearCache),
	}
}

func (c *Client) GetLocation(url *string) (Locations, error) {
	response, err := c.httpClient.Get(*url)
	if err != nil {
		return Locations{}, err
	}
	if val, ok := c.cache.Get(*url); ok {
		locationsResp := Locations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Locations{}, err
		}
		return locationsResp, nil
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		return Locations{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", response.StatusCode, body)
	}
	if err != nil {
		return Locations{}, err
	}
	locations := Locations{}
	err_json := json.Unmarshal(body, &locations)
	if err_json != nil {
		return Locations{}, err_json
	}
	c.cache.Add(*url, body)
	return locations, nil
}

func (c *Client) GetPokemonsInArea(name string) (LocationsArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	response, err := c.httpClient.Get(url)
	if err != nil {
		return LocationsArea{}, err
	}
	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationsArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationsArea{}, err
		}
		return locationsResp, nil
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		return LocationsArea{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", response.StatusCode, body)
	}
	if err != nil {
		return LocationsArea{}, err
	}
	locations := LocationsArea{}
	err_json := json.Unmarshal(body, &locations)
	if err_json != nil {
		return LocationsArea{}, err_json
	}
	c.cache.Add(url, body)
	return locations, nil
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	response, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", response.StatusCode, body)
	}
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err_json := json.Unmarshal(body, &pokemon)
	if err_json != nil {
		return Pokemon{}, err_json
	}

	c.cache.Add(url, body)
	return pokemon, nil
}
