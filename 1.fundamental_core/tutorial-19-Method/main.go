package main

import "fmt"

// Make a Struct
type Person struct {
	Name string
	Age int
}

// Make a Method Value Receiver
func (p Person) Greet()  {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}


// Make a Method Pointer Receiver
func (p *Person) HaveBirthday()  {
	p.Age++
}


func main()  {
	person := Person{Name: "Alice", Age: 30}
	// Before
    person.Greet()

	// After
    person.HaveBirthday()
    person.Greet()
}

// Note :
// When using Value Receiver or Pointer Receiver??

// Value Receiver:
// 1. If the method does not change the data in the struct.
// 2. If the struct is small and copying the struct is more efficient.

// Pointer Receiver:
// 1. If the method changes data in the struct.
// 2. If the struct is large and you want to avoid copy overhead.