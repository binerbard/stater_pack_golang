package main

import "fmt"

// fact calculates the factorial of a given integer 'n'.
func fact(n int) int {
    // Base case: If 'n' is 0, return 1 (by definition, 0! = 1).
    if n == 0 {
        return 1
    }
    // Recursive case: Multiply 'n' with the factorial of (n-1).
    return n * fact(n-1)
}

func main() {
    // Calculate and print the factorial of 7 using the 'fact' function.
    fmt.Println("Factorial of 7:", fact(7))

    // Declare a variable 'fib' of function type that computes Fibonacci numbers.
    var fib func(n int) int

    // Assign a closure function to 'fib' to compute Fibonacci numbers recursively.
    fib = func(n int) int {
        // Base case: If 'n' is less than 2, return 'n'.
        if n < 2 {
            return n
        }
        // Recursive case: Compute the (n-1)th Fibonacci number plus the (n-2)th Fibonacci number.
        return fib(n-1) + fib(n-2)
    }

    // Calculate and print the 7th Fibonacci number using the 'fib' function.
    fmt.Println("7th Fibonacci number:", fib(7))
}
