package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jayrgarg/pokedexcli/internal/config"
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

    initial_next :=  "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

    cfg := &config.Config{
        Next: &initial_next,
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
