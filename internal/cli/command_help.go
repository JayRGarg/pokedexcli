package cli

import (
    "fmt"
    "github.com/jayrgarg/pokedexcli/internal/config"
)

func commandHelp(cfg *config.Config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, v := range GetCmdMap() {
        fmt.Println(v.Name+":", v.Description)
    }
    return nil
}
