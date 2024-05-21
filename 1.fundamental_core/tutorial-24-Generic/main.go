// Generic is available on Golang 1.18+
// Generic is to make code more reusable and type safe

package main

import "fmt"

// 1. Function Generic
// This case is to write function with any tipe data
// for example : Max -> A, B

func Similarity[T comparable](a, b T) string {
	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)	
	if a == b {
		return aStr + " is same from " + bStr
	}
	return aStr + " is diferent from " + bStr
}


// 2. Struct Generic
//  This case is to deifne Structure with Any Type
//  for example : Pair
type Pair[T any,U any] struct {
	First T
	Seccond U
}

// 3. Interfaace Generic
// This case is to define motode with any tipe data
// for example : Basic Encapsulation
type Containaer[T any] interface{
	Get() T
	Set(T)
}

type Box [T any] struct {
	Value T
}

//  Why using Pointer? "*" is a pointer
// Because the value of variable is posible to changed
func (b *Box[T]) Get() T {
	return b.Value
}

func (b *Box[T]) Set(value T) {
	b.Value = value
}


func main()  {

	// 1. Function Generic Execution
	fmt.Println("1. Function Generic Example")
	fmt.Println(Similarity(3,5))
	fmt.Println(Similarity(5,5))
	fmt.Println(Similarity("apple","banana"))
	fmt.Println(Similarity("apple","apple"))
	
	
	// 2. Struct Generic Execution
	fmt.Println("\n2. Struct Generic Example")
	numberStringPair := Pair[int,string]{1,"one"}
	fmt.Printf("Pair: %v\n",numberStringPair)
	floatBoolPair := Pair[float64, bool]{3.14,true}
	fmt.Printf("Pair: %v\n", floatBoolPair)

	// 3. Interface Generic Execution
	fmt.Println("\n3. Interface Generic Example")
	// Initialize Box
	intBox := &Box[int]{}
	// Set Value
	intBox.Set(99)
	// Get Value
	fmt.Println("The value of intBox is:",intBox.Get())

	// Initialize Box
	stringBox := &Box[string]{}
	// Set Value
	stringBox.Set("Soul Game")
	// Get Value
	fmt.Println("The value of stringBox is:",stringBox.Get())
}