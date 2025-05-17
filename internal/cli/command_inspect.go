package cli

import (
    "fmt"
    "github.com/jayrgarg/pokedexcli/internal/config"
)

func commandInspect(cfg *config.Config, params *[]string) error {

    if (params == nil || len(*params) != 1) {
        return fmt.Errorf("Expecting 1 argument: pokemonName")
    }

    pokemonName := (*params)[0]
    pokemon, exists := (*cfg.PokeApiClient.Pokedex)[pokemonName]
    if !exists {
        return fmt.Errorf("you have not caught that pokemon")
    }

    fmt.Printf("Name: %v\n", pokemon.Name)
    fmt.Printf("Height: %v\n", pokemon.Height)
    fmt.Printf("Weight: %v\n", pokemon.Weight)

    fmt.Printf("Stats: \n")
    for _, stat := range pokemon.Stats {
        fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Printf("Types: \n")
    for _, t := range pokemon.Types {
        fmt.Printf(" - %v\n", t.Type.Name)
    }

    return nil
}
