package cli

import (
    "fmt"
    "sort"
    "github.com/jayrgarg/pokedexcli/internal/config"
)

func commandPokedex(cfg *config.Config, params *[]string) error {

    numPokemon := len(*cfg.PokeApiClient.Pokedex)
    if numPokemon == 0 {
        fmt.Println("Your Pokedex is Empty :(")
        return nil
    }

    keys := make([]string, 0, numPokemon)
    for k := range *cfg.PokeApiClient.Pokedex {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    fmt.Println("Your Pokedex :)")
    for _, k := range keys {
        fmt.Printf(" - %v\n", k)
    }
    return nil
}

