package cli

import (
    "github.com/jayrgarg/pokedexcli/internal/config"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
}


func GetCmdMap() map[string]CliCommand {
    return map[string]CliCommand {
        "help": {
            Name:        "help",
            Description: "Displays a help message",
            Callback:    commandHelp,
        },
        "exit": {
            Name:        "exit",
            Description: "Exit the Pokedex",
            Callback:    commandExit,
        },
        "map": {
            Name:        "map",
            Description: "Displays the next 20 locations",
            Callback:    commandMap,
        },
        "mapb": {
            Name:        "mapb",
            Description: "Displays the previous 20 locations",
            Callback:    commandMapB,
        },
    }

}
