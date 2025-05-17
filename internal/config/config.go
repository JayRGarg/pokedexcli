package config

import (
    "github.com/jayrgarg/pokedexcli/internal/pokeapi"
)

type Config struct {
    PokeApiClient       pokeapi.Client
    Next                *string
    Previous            *string
}
