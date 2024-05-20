package main

import (
	"fmt"
	"math"
)

// Interface
type geometry interface{
	area() float64
}


// Make Struct to implement interface in ractangle
type rectangle struct{
	width, height float64
}

// Make Struct to implement interface in circle
type circle struct{
	radius float64
}


// Method Area for rectangle
func (r rectangle) area() float64{
	return r.width * r.height
}


// Method Area for circle
func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}


// Implementing interface to treat Method input as the Purpose
func formula(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main()  {
	
	// Define rectangle struct value
	square := rectangle{width: 3, height: 5}

	// Define circle struct value
	circle := circle{ radius: 4}


	// For square
	fmt.Println("For a rectangle")
	formula(square)


	// For Circle
	fmt.Println("\nFor a circle")
	formula(circle)

}

