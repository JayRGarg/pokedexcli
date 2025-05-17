package cli

import (
    "github.com/jayrgarg/pokedexcli/internal/config"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, *[]string) error
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
        "explore": {
            Name:        "explore",
            Description: "Explores a location area - takes 1 parameter: locationName",
            Callback:    commandExplore,
        },
        "catch": {
            Name:        "catch",
            Description: "Attempts to catch a Pokemon - takes 1 parameter: pokemonName",
            Callback:    commandCatch,
        },
        "inspect": {
            Name:        "inspect",
            Description: "Inspects a caught Pokemon - takes 1 parameter: pokemonName",
            Callback:    commandInspect,
        },
        "pokedex": {
            Name:        "pokedex",
            Description: "Displays the list of Pokemon caught in your Pokedex",
            Callback:    commandPokedex,
        },
    }

}
