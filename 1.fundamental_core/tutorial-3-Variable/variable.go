package main

import (
	"fmt"
)

func main()  {
	
	var a int = 1
	var b int = 2
	var c int = a + b
	fmt.Println(c)

	var d int
	var e int
	d = 1
	e = 2
	fmt.Println(d + e)

	var name string = "example 1"
	fmt.Println(name)

	var name2 = "example 2"
	fmt.Println(name2)

	name3 := "example 3"
	fmt.Println(name3)

	number:= 1
	fmt.Println(number)

	number = 2
	fmt.Println(number)

	var (
		name4 string = "example 4"
		name5 string = "example 5"
	)
	fmt.Println(name4)
	fmt.Println(name5)

	var name6, name7 = "example 6", "example 7"
	fmt.Println(name6)
	fmt.Println(name7)

	var biner bool = true
	fmt.Println(biner)

	var biner2 = false
	fmt.Println(biner2)

	var (
		biner3 bool = true
		biner4 bool = false
	)
	fmt.Println(biner3)
	fmt.Println(biner4)

	biner5 := true
	biner6 := false
	fmt.Println(biner5)
	fmt.Println(biner6)
	fmt.Println(!biner6)

	var decimal float32 = 1.1 // 32 bit float number 
	fmt.Println(decimal)

	var decimal2 float64 = 1.2
	fmt.Println(decimal2)

}