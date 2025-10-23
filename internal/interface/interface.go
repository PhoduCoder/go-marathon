package main

import "fmt"

// Defining an interface
type Shape interface {
	Name() string
	Area() float64
}

//Define three structs for triangle, circle and square

type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	width  float64
	length float64
}

type circle struct {
	radius float64
}

func (t triangle) Name() string {
	return "Triangular Shape"
}

func (t triangle) Area() float64 {
	return (0.5 * t.base * t.height)
}

func (r rectangle) Name() string {
	return "Rectangular Shape"
}

func (r rectangle) Area() float64 {
	return (0.5 * r.length * r.width)
}

func (c circle) Name() string {
	return "Circular Shape"
}

func (c circle) Area() float64 {
	return (3.14 * c.radius * c.radius)
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	c := circle{radius: 10}
	printShapeDetails(t, r, c)
}
