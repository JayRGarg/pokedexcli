package cli

import (
    "fmt"
    "sort"
    "github.com/jayrgarg/pokedexcli/internal/config"
)

func commandHelp(cfg *config.Config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")

    commands := GetCmdMap()
    keys := make([]string, 0, len(commands))
	for name := range commands {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	for _, name := range keys {
        fmt.Println(commands[name].Name+":", commands[name].Description)
	}
    // for _, v := range GetCmdMap() {
    //     fmt.Println(v.Name+":", v.Description)
    // }
    return nil
}
