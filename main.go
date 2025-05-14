package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func cleanInput(text string) []string {
    clean := strings.Fields(text)
    for i := range clean {
        clean[i] = strings.ToLower(clean[i])
    }    
    return clean
}

func commandExit(cfg *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, v := range cmdMap {
        fmt.Println(v.name+":", v.description)
    }
    return nil
}

func commandMap(cfg *config) error {
    url := cfg.Next
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return fmt.Errorf("Error formatting request: %w", err)
    }

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Error performing request from Pokedox API: %w", err)
    }
    defer res.Body.Close()

    resourcesBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return fmt.Errorf("Error reading response to bytes: %w", err)
    }

    var resourcesParsed Resources
    err = json.Unmarshal(resourcesBytes, &resourcesParsed)
    if err != nil {
        return fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    for _, v := range resourcesParsed.Results {
        fmt.Println(v.Name)
    }

    cfg.Next = resourcesParsed.Next
    cfg.Previous = resourcesParsed.Previous

    return nil
}

func commandMapB(cfg *config) error {
    url := cfg.Previous
    if url == "" {
        if strings.Contains(cfg.Next, "offset=0") {
            fmt.Println("Currently haven't gone to First Page")
            return nil
        } else if strings.Contains(cfg.Next, "offset=20") {
            fmt.Println("Currently on First Page, go to any of the next pages with the 'map' command")
            return nil
        } else {
            return fmt.Errorf("Error: cfg.Previous empty with unexpected cfg.Next")
        }
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return fmt.Errorf("Error formatting request: %w", err)
    }

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Error performing request from Pokedox API: %w", err)
    }
    defer res.Body.Close()

    resourcesBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return fmt.Errorf("Error reading response to bytes: %w", err)
    }

    var resourcesParsed Resources
    err = json.Unmarshal(resourcesBytes, &resourcesParsed)
    if err != nil {
        return fmt.Errorf("Error Unmarshalling resources to struct: %w", err)
    }

    for _, v := range resourcesParsed.Results {
        fmt.Println(v.Name)
    }

    cfg.Next = resourcesParsed.Next
    cfg.Previous = resourcesParsed.Previous

    return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
    Next        string
    Previous    string
}

type Resource struct {
    Name        string
    Url         string
}

type Resources struct {
    Count       int
    Next	    string
    Previous	string
    Results     []Resource
}

var cmdMap map[string]cliCommand

func main() {

    cfg := &config{
        Next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
        Previous: "",
    }

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
        "map": {
            name:        "map",
            description: "Displays the next 20 locations",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the previous 20 locations",
            callback:    commandMapB,
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
            err := cmd.callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}
