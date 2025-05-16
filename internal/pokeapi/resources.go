package pokeapi

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
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

func GetLocationAreaResources (url *string) (Resources, error) {
    if url == nil {
        return Resources{}, fmt.Errorf("Error: URL is a nil pointer")
    }
    req, err := http.NewRequest("GET", *url, nil)
    if err != nil {
        return Resources{}, fmt.Errorf("Error formatting request: %w", err)
    }

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return Resources{}, fmt.Errorf("Error performing request from Pokedox API: %w", err)
    }
    defer res.Body.Close()

    resourcesBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return Resources{}, fmt.Errorf("Error reading response to bytes: %w", err)
    }

    var resourcesParsed Resources
    err = json.Unmarshal(resourcesBytes, &resourcesParsed)
    if err != nil {
        return Resources{}, fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    return resourcesParsed, nil

}
