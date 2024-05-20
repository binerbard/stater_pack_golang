package main

import (
	"fmt"
)

func main() {
    // Declare an uninitialized slice 's' of strings
    var s []string
    // Print the uninitialized slice, check if it's nil, and if its length is 0
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    // Initialize a slice 's' with a length of 3 using the make function
    s = make([]string, 3)
    // Print the initialized slice along with its length and capacity
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    // Set values at specific indices in slice 's'
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    // Print the slice after setting values and retrieve the value at index 2
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // Print the length of slice 's'
    fmt.Println("len:", len(s))

    // Append elements to slice 's'
    s = append(s, "d")
    s = append(s, "e", "f")
    // Print the slice after appending
    fmt.Println("apd:", s)

    // Create a new slice 'c' with the same length as slice 's' and copy values from 's' to 'c'
    c := make([]string, len(s))
    copy(c, s)
    // Print the copied slice 'c'
    fmt.Println("cpy:", c)

    // Slice 's' to create a new slice 'l' from indices 2 to 4
    l := s[2:5]
    fmt.Println("sl1:", l)

    // Slice 's' to create a new slice 'l' from the beginning to index 4
    l = s[:5]
    fmt.Println("sl2:", l)

    // Slice 's' to create a new slice 'l' from index 2 to the end
    l = s[2:]
    fmt.Println("sl3:", l)

    // Declare and initialize a new slice 't' directly
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Create a two-dimensional slice 'twoD' with 3 inner slices
    twoD := make([][]int, 3)
    // Initialize each inner slice with varying lengths and populate with values based on indices
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    // Print the two-dimensional slice
    fmt.Println("2d: ", twoD)
}
