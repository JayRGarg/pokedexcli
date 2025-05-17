package pokeapi

import (
    "net/http"
    "time"
	"github.com/jayrgarg/pokedexcli/internal/pokecache"
)

type Pokemon struct {
    Name            string
    Height          int
    Weight          int
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func NewPokedex() *map[string]Pokemon {
    pokedex := make(map[string]Pokemon)
    return &pokedex
}

type Client struct {
    httpClient      *http.Client
    cache           *pokecache.Cache
    Pokedex         *map[string]Pokemon
}

func NewClient(clientTimeout, cacheInterval time.Duration) *Client {
    return &Client{
        httpClient: &http.Client{
            Timeout: clientTimeout,
        },
        cache:      pokecache.NewCache(cacheInterval),
        Pokedex:    NewPokedex(),
    }
}
