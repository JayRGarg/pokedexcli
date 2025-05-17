package main

import (
    "testing"
    "os"
    "bytes"
    "io"
    "strings"
    "time"
	"github.com/jayrgarg/pokedexcli/internal/config"
	"github.com/jayrgarg/pokedexcli/internal/pokeapi"
	"github.com/jayrgarg/pokedexcli/internal/cli"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "  hello jay's world  ",
            expected: []string{"hello", "jay's", "world"},
        },
        {
            input:    "hello jay's world  ",
            expected: []string{"hello", "jay's", "world"},
        }, 
        {
            input:    "  hello jay's world",
            expected: []string{"hello", "jay's", "world"},
        }, 
        {
            input:    "hello jay's world",
            expected: []string{"hello", "jay's", "world"},
        }, 
        {
            input:    "hellojay'sworld",
            expected: []string{"hellojay'sworld"},
        }, 
        // add more cases here
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Length of actual output [%d] and expected output[%d] do not match!", len(actual), len(c.expected))
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("First mismatch detected at index: %d, actual '%s' and expected '%s' do not match!", i, word, expectedWord)
            }
        }
    }
}

func TestCatchCommand(t *testing.T) {
    // Save the original stdout
    oldStdout := os.Stdout

    for range 10 {
        // Create a pipe
        r, w, _ := os.Pipe()
        os.Stdout = w
        
        cfg := &config.Config{
            PokeApiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
            //PokeApiClient: pokeapi.NewClient(5 * time.Second, 10 * time.Second),
            Next: nil,
            Previous: nil,
        }

        cmdMap := cli.GetCmdMap()
        cmd := cmdMap["catch"]
        args := &[]string{"tentacool"}
        err :=  cmd.Callback(cfg, args)
        if err != nil {
            t.Errorf("Error in calling 'catch': %v", err)
        }
        
        // Close the write end of the pipe to flush it
        w.Close()
        
        // Restore the original stdout
        os.Stdout = oldStdout
        
        // Read the captured output
        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()
        
        if strings.Contains(output, "caught") {
            t.Log("Caught!")
            pokemon, exists := (*cfg.PokeApiClient.Pokedex)["tentacool"]
            name := "tentacool"
            expected := pokeapi.Pokemon{ Name: name}
            if !exists {
                t.Errorf("Caught but not added to pokedex")
            }
            if pokemon.Name != expected.Name {
                t.Errorf("Pokemon in Pokedex does not match expected")
            }
        } else {
            t.Log("Escaped!")
            _, exists := (*cfg.PokeApiClient.Pokedex)["tentacool"]
            if exists {
                t.Errorf("Not Caught but was still added to pokedex")
            }
        }
    }
}
