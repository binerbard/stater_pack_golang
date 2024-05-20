package main

import "fmt"

// intSeq returns a closure function that generates sequential integers starting from 1.
func intSeq() func() int {
    // Initialize a variable 'i' with value 0.
    i := 0

    // Return a closure function.
    return func() int {
        // Increment the value of 'i' by 1.
        i++
        // Return the current value of 'i'.
        return i
    }
}

func main() {
    // Call intSeq to create a new closure function 'nextInt'.
    // This closure function captures the variable 'i' from intSeq.
    nextInt := intSeq()

    // Call 'nextInt' three times and print the generated integers.
    fmt.Println(nextInt()) // Output: 1
    fmt.Println(nextInt()) // Output: 2
    fmt.Println(nextInt()) // Output: 3

    // Call intSeq again to create a new closure function 'newInts'.
    // This time, a new variable 'i' is created because intSeq is called again.
    newInts := intSeq()

    // Call 'newInts' and print the generated integer.
    // Since it's a new closure, it starts counting from 1 again.
    fmt.Println(newInts()) // Output: 1
}
