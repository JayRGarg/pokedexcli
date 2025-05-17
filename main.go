package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "time"
	"github.com/jayrgarg/pokedexcli/internal/config"
	"github.com/jayrgarg/pokedexcli/internal/pokeapi"
	"github.com/jayrgarg/pokedexcli/internal/cli"
)

func cleanInput(text string) []string {
    clean := strings.Fields(text)
    for i := range clean {
        clean[i] = strings.ToLower(clean[i])
    }    
    return clean
}

func main() {

    cfg := &config.Config{
        PokeApiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
        //PokeApiClient: pokeapi.NewClient(5 * time.Second, 10 * time.Second),
        Next: nil,
        Previous: nil,
    }

    cmdMap := cli.GetCmdMap()


    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        inputTxt := scanner.Text()
        tokens := cleanInput(inputTxt)
        if len(tokens) == 0 {
            continue
        }
        cmd, exists := cmdMap[tokens[0]]
        if !exists {
            fmt.Println("Unknown command")
        } else {
            err := cmd.Callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}
