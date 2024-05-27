package main

import (
	"fmt"
	"math"
)

type Circle struct {
	name   string
	Radius float64
}

type Rectangle struct {
	name   string
	Width  float64
	Height float64
}

type Shape interface {
	area() float64 // Method with no arguments and returns a float64
	changeName(string)
}

func main() {
	circle := Circle{Radius: 4.0}
	rectangle := Rectangle{Width: 30.0, Height: 6.0}

	shapes := []Shape{&circle, &rectangle}

	for _, shape := range shapes {
		fmt.Println("area:", shape.area())
	}

}

// Binding the struct to the interface happens by implementing the method declared by the interface
// which is associated with a reference to the struct
func (c *Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) changeName(newName string) {
	c.name = newName
}

func (r *Rectangle) area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) changeName(newName string) {
	r.name = newName
}
