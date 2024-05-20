package main

import (
	"fmt"
)

func main() {
    // Create an empty map 'm' with keys of type string and values of type int
    m := make(map[string]int)

    // Set key-value pairs in the map 'm'
    m["k1"] = 7
    m["k2"] = 13

    // Print the entire map 'm'
    fmt.Println("map:", m)

    // Access the value associated with key "k1" in map 'm'
    v1 := m["k1"]
    fmt.Println("v1:", v1)

    // Access the value associated with key "k3" in map 'm' (which does not exist)
    // The value retrieved for a non-existent key is the zero value for the value type
    v3 := m["k3"]
    fmt.Println("v3:", v3)

    // Print the number of key-value pairs in map 'm'
    fmt.Println("len:", len(m))

    // Delete the key-value pair with key "k2" from map 'm'
    delete(m, "k2")
    fmt.Println("map:", m)

    // Check if key "k2" exists in map 'm' using the blank identifier to discard the value
    // 'prs' will be true if the key exists, false otherwise
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // Declare and initialize a new map 'n' with key-value pairs
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)


}
