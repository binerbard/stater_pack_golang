package main

import (
	"fmt"
	"time"
)

func f(from string) {

	// Generate code with 0 to 2
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
		time.Sleep(time.Second)
    }
}

func main() {

	// This is a direct couting
	fmt.Println("Direct Counting Start")
    f("direct")
	
	fmt.Println("Direct Counting Done")
	
	fmt.Println("\nGoroutines Counting Start")
    go f("goroutine")
	
	// Intruption with message "going"
    go func(msg string) {
		time.Sleep(time.Second)
		fmt.Println(msg)
		}("going")
	// Make a session 5 second before next counting
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutines Counting Done")


	// OUTPUT :
	// Direct Counting Start
	// direct : 0
	// direct : 1
	// direct : 2
	// Direct Counting Done

	// Goroutines Counting Start
	// goroutine : 0
	// goroutine : 1
	// going
	// goroutine : 2
}