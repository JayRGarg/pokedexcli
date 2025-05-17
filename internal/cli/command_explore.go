package cli

import (
    "fmt"
    "github.com/jayrgarg/pokedexcli/internal/config"
)

func commandExplore(cfg *config.Config, params *[]string) error {

    if (params == nil || len(*params) != 1) {
        return fmt.Errorf("Expecting 1 argument: locationName")
    }

    locationName := (*params)[0]
    locationArea, err := cfg.PokeApiClient.GetLocationAreaInfo(&locationName)
    
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %v...\n", locationName)
    fmt.Println("Found Pokemon:")
    for _, v := range locationArea.PokemonEncounters {
        fmt.Println(" - " + v.Pokemon.Name)
    }

    return nil
}

