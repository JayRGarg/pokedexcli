package pokeapi

import (
    "net/http"
    "time"
	"github.com/jayrgarg/pokedexcli/internal/pokecache"
)

type Client struct {
    httpClient      *http.Client
    cache           *pokecache.Cache
}

func NewClient(clientTimeout, cacheInterval time.Duration) *Client {
    return &Client{
        httpClient: &http.Client{
            Timeout: clientTimeout,
        },
        cache:      pokecache.NewCache(cacheInterval),
    }
}
