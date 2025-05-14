package main

import (
	"bufio"
    "os"
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
    clean := strings.Fields(text)
    for i := range clean {
        clean[i] = strings.ToLower(clean[i])
    }    
    return clean
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, v := range cmdMap {
        fmt.Println(v.name+":", v.description)
    }
    return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cmdMap map[string]cliCommand

func main() {

    cmdMap = map[string]cliCommand {
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
    }


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
            err := cmd.callback()
            if err != nil {
                fmt.Println("Error:", err)
            }
        }
    }
}
