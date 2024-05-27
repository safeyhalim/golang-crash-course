package main

import (
	"fmt"
	"math"

	"github.com/sixfwa/crashcourse/randomstuff"
	"github.com/sixfwa/crashcourse/shapes"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Pi)

	circle := shapes.Circle{Radius: 4.5}
	rectangle := shapes.Rectangle{Width: 10.0, Height: 13.0}
	fmt.Println("Area of circle:", circle.Area())
	fmt.Println("Area of  rectangle:", rectangle.Area())
	fmt.Println("Is rectangle a square", rectangle.IsSquare())

	randomstuff.PrintSomeRubbish()
}
