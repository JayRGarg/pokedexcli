package pokeapi

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Resources struct {
    Count       int         `json:"count"`
    Next	    *string      `json:"next"`
    Previous	*string      `json:"previous"`
    Results     []struct {
        Name    string      `json:"name"`
        Url     string      `json:"url"`
    } `json:"results"`
}

func (c *Client) GetLocationAreaResources (pageUrl *string) (Resources, error) {
    //tmpUrl := baseURL + "/location-area"//+"/?offset=0&limit=20"
    tmpUrl := baseURL + "/location-area"+"/?offset=0&limit=20"
    if pageUrl != nil {
        tmpUrl = *pageUrl
    }
    url := &tmpUrl

    var resourcesBytes []byte

    resourcesBytes, foundInCache := c.cache.Get(*url)

    if !foundInCache {
        req, err := http.NewRequest("GET", *url, nil)
        if err != nil {
            return Resources{}, fmt.Errorf("Error formatting request: %w", err)
        }

        res, err := c.httpClient.Do(req)
        if err != nil {
            return Resources{}, fmt.Errorf("Error performing request from Pokedox API: %w", err)
        }
        defer res.Body.Close()

        resourcesBytes, err = io.ReadAll(res.Body)
        if err != nil {
            return Resources{}, fmt.Errorf("Error reading response to bytes: %w", err)
        }

        c.cache.Add(*url, resourcesBytes)
    }


    var resourcesParsed Resources
    err := json.Unmarshal(resourcesBytes, &resourcesParsed)
    if err != nil {
        return Resources{}, fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    return resourcesParsed, nil

}
