package main

import (
	"fmt"
	"time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)
	c3 := make(chan string)

	// Running when second 1 (in Pararel)
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()

	go func() {
        time.Sleep(1 * time.Second)
        c3 <- "ping"
    }()

	// Running when second 2 (in Pararel)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 4; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
		case msg3 := <-c3:
            fmt.Println("received", msg3)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
        }
    }
}