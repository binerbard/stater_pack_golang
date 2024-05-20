package main

import (
	"fmt"
)

func main()  {
	fmt.Println("This language is "+"golang")
	fmt.Println("1+1",1+1)
	fmt.Println("2-1",2-1)
	fmt.Println("3*2",3*2)
	fmt.Println("4/2",4/2)
	fmt.Println("5%2",(5%2))
	fmt.Println("5&2",(5&2))
	fmt.Println("5|2",(5|2))
	fmt.Println("5^2",(5^2))
	fmt.Println("5<<2",(5<<2))
	fmt.Println("5>>2",(5>>2))

	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}