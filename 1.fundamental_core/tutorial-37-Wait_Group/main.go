// Wait Group
// to wait for all goroutines to finish before continuing execution

// sync.WaitGroup function:
// 1. Coordinates a group of goroutines.
// 2. Ensures that the program does not continue before all goroutines have finished.

// Ways of working:
// 1. The program creates a sync.WaitGroup.
// 2. Each goroutine to be run adds itself to the WaitGroup.
// 3. Each time a goroutine completes its task, it subtracts itself from the WaitGroup.
// 4. The program calls .Wait() on the WaitGroup to wait for all goroutines to complete.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Reduces the WaitGroup after the goroutine completes

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup // Create WaitGroup

    // Add goroutine to WaitGroup
    for i := 1; i <= 2; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    // Waiting all done
    wg.Wait()
    fmt.Println("All workers have finished")
}
