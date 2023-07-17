package main

import (
	"fmt"
)

// this  is a function in go lang
func add(a, b int) int {
	return a + b
}

type Rectangle struct {
	Width  float64
	Height float64
}

// method for type rectangle here the width and height must be be used in the method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
    //fn result
	result := add(3, 4)
	fmt.Println("Result:", result)  
	
	//mthod result
	rect := Rectangle{Width: 5, Height: 3} //store struct to rect.
	area := rect.Area()
	fmt.Println("Area:", area)
	
}

