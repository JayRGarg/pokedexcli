package main

import (
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

func main() {
    fmt.Println("Hello, World!")
}
