package main

import "fmt"

func main() {
    // Declare an array 'a' of length 5 initialized with zero values
    var a [5]int
    fmt.Println("emp:", a)

    // Set the value at index 4 of array 'a' to 100
    a[4] = 100
    fmt.Println("set:", a)
    
    // Get and print the value at index 4 of array 'a'
    fmt.Println("get:", a[4])

    // Print the length of array 'a'
    fmt.Println("len:", len(a))

    // Declare and initialize an array 'b' with values
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Redefine and initialize array 'b' with ellipsis
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Redefine and initialize array 'b' with specific index values
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

    // Declare a 2D array 'twoD' with 2 rows and 3 columns
    var twoD [2][3]int

    // Loop through 'twoD' to set values based on indices
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    // Redefine and initialize 'twoD' directly with values
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}
