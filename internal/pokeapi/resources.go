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

func (c *Client) GetLocationAreasResources (pageUrl *string) (LocationAreasResources, error) {
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
            return LocationAreasResources{}, fmt.Errorf("Error formatting request: %w", err)
        }

        res, err := c.httpClient.Do(req)
        if err != nil {
            return LocationAreasResources{}, fmt.Errorf("Error performing request from Pokedox API: %w", err)
        }
        defer res.Body.Close()

        resourcesBytes, err = io.ReadAll(res.Body)
        if err != nil {
            return LocationAreasResources{}, fmt.Errorf("Error reading response to bytes: %w", err)
        }

        c.cache.Add(*url, resourcesBytes)
    }


    var resourcesParsed LocationAreasResources
    err := json.Unmarshal(resourcesBytes, &resourcesParsed)
    if err != nil {
        return LocationAreasResources{}, fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    return resourcesParsed, nil

}

func (c *Client) GetLocationAreaInfo (name *string) (LocationAreaInfo, error) {
    if name == nil {
        return LocationAreaInfo{}, fmt.Errorf("Location Name Pointer is Nil!")
    }
    tmpUrl := baseURL + "/location-area/" + *name
    url := &tmpUrl

    var resourceBytes []byte

    resourceBytes, foundInCache := c.cache.Get(*url)

    if !foundInCache {
        req, err := http.NewRequest("GET", *url, nil)
        if err != nil {
            return LocationAreaInfo{}, fmt.Errorf("Error formatting request: %w", err)
        }

        res, err := c.httpClient.Do(req)
        if err != nil {
            return LocationAreaInfo{}, fmt.Errorf("Error performing request from Pokedox API: %w", err)
        }
        defer res.Body.Close()

        resourceBytes, err = io.ReadAll(res.Body)
        if err != nil {
            return LocationAreaInfo{}, fmt.Errorf("Error reading response to bytes: %w", err)
        }

        c.cache.Add(*url, resourceBytes)
    }


    var resourceParsed LocationAreaInfo 
    err := json.Unmarshal(resourceBytes, &resourceParsed)
    if err != nil {
        return LocationAreaInfo{}, fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    return resourceParsed, nil

}

func (c *Client) GetPokemonInfo (name *string) (PokemonInfo, error) {
    if name == nil {
        return PokemonInfo{}, fmt.Errorf("Pokemon Name Pointer is Nil!")
    }
    tmpUrl := baseURL + "/pokemon/" + *name
    url := &tmpUrl

    var resourceBytes []byte

    resourceBytes, foundInCache := c.cache.Get(*url)

    if !foundInCache {
        req, err := http.NewRequest("GET", *url, nil)
        if err != nil {
            return PokemonInfo{}, fmt.Errorf("Error formatting request: %w", err)
        }

        res, err := c.httpClient.Do(req)
        if err != nil {
            return PokemonInfo{}, fmt.Errorf("Error performing request from Pokedox API: %w", err)
        }
        defer res.Body.Close()

        resourceBytes, err = io.ReadAll(res.Body)
        if err != nil {
            return PokemonInfo{}, fmt.Errorf("Error reading response to bytes: %w", err)
        }

        c.cache.Add(*url, resourceBytes)
    }


    var resourceParsed PokemonInfo 
    err := json.Unmarshal(resourceBytes, &resourceParsed)
    if err != nil {
        return PokemonInfo{}, fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    return resourceParsed, nil

}
