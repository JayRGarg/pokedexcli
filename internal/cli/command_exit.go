package cli

import (
    "fmt"
    "github.com/jayrgarg/pokedexcli/internal/config"
    "os"
)

func commandExit(cfg *config.Config, params *[]string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}
