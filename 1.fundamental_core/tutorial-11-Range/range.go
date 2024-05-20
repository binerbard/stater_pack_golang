package main

import "fmt"

func main() {
    // Create a slice of integers
    nums := []int{2, 3, 4}
    sum := 0

    // Calculate the sum of all numbers in the slice using a for loop
    for _, num := range nums {
        sum += num
    }
    // Print the sum
    fmt.Println("sum:", sum)

    // Iterate over the slice 'nums' using range and print the index of the element with value 3
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // Create a map 'kvs' with string keys and string values
    kvs := map[string]string{"a": "apple", "b": "banana"}

    // Iterate over the key-value pairs in the map 'kvs' using range and print each pair
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // Iterate over the keys in the map 'kvs' using range and print each key
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // Iterate over the string "go" using range and print the index and ASCII value of each character
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
