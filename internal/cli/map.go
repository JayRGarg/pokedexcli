package cli

import (
    "fmt"
    //"strings"
    "errors"
    "github.com/jayrgarg/pokedexcli/internal/config"
)

func commandMap(cfg *config.Config) error {
    url := cfg.Next

    locationAreaResouces, err := cfg.PokeApiClient.GetLocationAreaResources(url)
    
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
        if cfg.Next == nil {// || strings.Contains(*cfg.Next, "offset=0") {
            return errors.New("Currently haven't gone to First Page")
        } else {//if strings.Contains(*cfg.Next, "offset=20") {
            return errors.New("Currently on First Page, go to any of the next pages with the 'map' command")
        }
        //may need to add logic here? not sure if this case can exist though
    }

    locationAreaResouces, err := cfg.PokeApiClient.GetLocationAreaResources(url)
    
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
