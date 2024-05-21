// The term "lexicographic order" refers to the order of words or strings
// based on the alphabetical order of their characters.

package main

import (
	"fmt"
)

// Define a constraint for types that support condition graterthan and lessthan (<,>,<=,>=)
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Max("apple", "banana")) // Output : Banana
}

// Explain:
// Apple have fist character is "A" or in Unicode Value is (A = 97 -> Unicode)
// Banana have fist character is "B" or in Unicode Value is (B = 98 -> Unicode)

// Why the answer is Banana ??
// From Apple > Banana, is look like 97 > 98 in unicode