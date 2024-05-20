package main

import "fmt"

// zeroval takes an integer value and sets it to zero.
// Since it takes a copy of the value, the original value remains unchanged.
func zeroval(ival int) {
    ival = 0
}

// zeroptr takes a pointer to an integer and sets the value at that address to zero.
// This modifies the original value since it works with the pointer to the value.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    // Call zeroval with the integer value.
    zeroval(i)
    fmt.Println("zeroval:", i)  // The original value of i is unchanged.

    // Call zeroptr with the address of i.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)  // The original value of i is changed to 0.

    // Print the memory address of i.
    fmt.Println("pointer:", &i)
}
