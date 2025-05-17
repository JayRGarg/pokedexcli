package cli

import (
    "fmt"
    "math/rand"
    "github.com/jayrgarg/pokedexcli/internal/config"
    "github.com/jayrgarg/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config.Config, params *[]string) error {

    if (params == nil || len(*params) != 1) {
        return fmt.Errorf("Expecting 1 argument: pokemonName")
    }

    pokemonName := (*params)[0]
    pokemonInfo, err := cfg.PokeApiClient.GetPokemonInfo(&pokemonName)
    
    if err != nil {
        return err
    }

    fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
    client_roll := rand.Float32() * 100
    base_exp := pokemonInfo.BaseExperience
    fmt.Printf("pokemon base exp: %v\n", base_exp)
    fmt.Printf("client roll: %v\n", client_roll)
    if client_roll > float32(base_exp) {
        (*cfg.PokeApiClient.Pokedex)[pokemonName] = pokeapi.Pokemon{ 
                                                                        Name: pokemonName,
                                                                        Height: pokemonInfo.Height,
                                                                        Weight: pokemonInfo.Weight,
                                                                        Stats: pokemonInfo.Stats,
                                                                        Types: pokemonInfo.Types,
                                                                    }
        fmt.Printf("%v was caught!\n", pokemonName)
    } else {
        fmt.Printf("%v escaped!\n", pokemonName)
    }

    return nil
}

