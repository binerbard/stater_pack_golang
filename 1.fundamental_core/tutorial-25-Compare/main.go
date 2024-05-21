// Compare is a contraint type data for compare
// This purpose is to compare data like "==" or "!="
package main

import "fmt"

func Checking[T comparable] (a, b T) string {
	if a == b {
		return fmt.Sprintf("%v is same from %v", a, b)
	}
	return fmt.Sprintf("%v is different from %v", a, b)
}

func main()  {
	fmt.Println(Checking(1,2))
}