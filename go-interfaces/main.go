package main

import "fmt"

type shape interface {
	area() float64
	circum() float64
}

type square struct {
	sideLength float64
}

type circle struct {
	radius float64
}

//? square methods
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (s square) circum() float64 {
	return s.sideLength * 4
}

//? circle methods
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) circum() float64 {
	return 2 * 3.14 * c.radius
}

func printShapeInfo(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Circumference:", s.circum())
}

func main() {

	//? slice of shapes
	var shapes = []shape{
		square{sideLength: 5.0},
		circle{radius: 2.0},
	}

	for _, shape := range shapes {
		printShapeInfo(shape)
		fmt.Println("---")
	}

}
