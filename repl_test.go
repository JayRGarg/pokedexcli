package main

import "testing"

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
