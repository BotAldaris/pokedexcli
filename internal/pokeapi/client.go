package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocation(url *string) (Locations, error) {
	response, err := c.httpClient.Get(*url)
	if err != nil {
		return Locations{}, err
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
	return locations, nil
}
