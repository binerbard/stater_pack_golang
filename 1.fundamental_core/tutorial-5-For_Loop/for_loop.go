package main

import "fmt"

func main() {

	// Normal For Loop

	fmt.Println("Normal For Loop")

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    for i := 0; i < 3; i++ {
        fmt.Println("range", i)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n < 6; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }


	// Iteration Slice Array
	fmt.Println("Slice Array")
	nums := []int{2, 3, 4}
    for i, num := range nums {
        fmt.Println("index:", i, "value:", num)
    }
	

	// Iteration Map
	fmt.Println("Map")
	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Println("key:", k, "value:", v)
    }

	// Iteration Character
	fmt.Println("Character")
	str := "hello"
    for i, c := range str {
        fmt.Println("index:", i, "character:", string(c))
    }

	// Iteration by Channel
	fmt.Println("Channel")
	// Example 1
	fmt.Println("Example 1")
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

	// Example 2
	fmt.Println("Example 2")
	cha := make(chan int, 2)
    cha <- 1
    cha <- 2
    close(cha)

    for v := range cha {
        fmt.Println(v)
    }

	// Simulation Range With number
	fmt.Println("Simulation Range With number")
	// Make slice from 0 to 5
	numbers := []int{0, 1, 2, 3, 4, 5}

	// Use range to iterate the slice
	for _, num := range numbers {
		fmt.Println(num)
	}


	// If you want use range to generate number
	fmt.Println("If you want use range to generate number")
	for i := range make([]struct{}, 10) {
        fmt.Println(i)
    }
}
