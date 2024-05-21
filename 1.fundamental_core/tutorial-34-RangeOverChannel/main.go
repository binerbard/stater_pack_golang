package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)


	// Using range to fetch data from channel queue
    for elem := range queue {
        fmt.Println(elem)
    }
}