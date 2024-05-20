package main

import (
	"fmt"
	"math"
)

const s string = "constant"
func main() {

	const PI = 3.14

	fmt.Println("PI1 ",PI)

	const PI2 = PI

	fmt.Println("PI2 ",PI2)

	const PI3 = math.Pi

	fmt.Println("PI3 ",PI3)

	const divide = PI3 / 2
	fmt.Println("divide ",divide)

	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n // 3e20 = 300000000000000000000
    fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}