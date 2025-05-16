package cli

import (
    "fmt"
    "strings"
    "github.com/jayrgarg/pokedexcli/internal/config"
    "github.com/jayrgarg/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config.Config) error {
    url := cfg.Next

    locationAreaResouces, err := pokeapi.GetLocationAreaResources(url)
    
    if err != nil {
        return err
    }

    for _, v := range locationAreaResouces.Results {
        fmt.Println(v.Name)
    }

    cfg.Next = locationAreaResouces.Next
    cfg.Previous = locationAreaResouces.Previous

    return nil
}

func commandMapB(cfg *config.Config) error {
    url := cfg.Previous
    if url == nil {
        if strings.Contains(*cfg.Next, "offset=0") {
            fmt.Println("Currently haven't gone to First Page")
            return nil
        } else if strings.Contains(*cfg.Next, "offset=20") {
            fmt.Println("Currently on First Page, go to any of the next pages with the 'map' command")
            return nil
        }
    }

    locationAreaResouces, err := pokeapi.GetLocationAreaResources(url)
    
    if err != nil {
        return err
    }

    for _, v := range locationAreaResouces.Results {
        fmt.Println(v.Name)
    }

    cfg.Next = locationAreaResouces.Next
    cfg.Previous = locationAreaResouces.Previous

    return nil
}
