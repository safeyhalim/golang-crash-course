package main

import "fmt"

func main() {
	fmt.Println("The average:", calculateAverage(5.5, 4.5, 6.7, 9.8))
	boringFunction("Nice world", 324.9, "covid sucks", 3234)

	// Internal function (function inside of a function)
	func() {
		fmt.Println("My internal function")
	}()

	// Internal function assigned to a variable and being called later
	add := func(numberOne, numberTwo int) int {
		return numberOne + numberTwo
	}

	// Calling the internal function
	fmt.Println("Using the internal add():", add(34, 5))

	// Clojure: Internal function that uses a variable outside its scope
	hundred := 100
	addHundred := func() int {
		hundred += 100
		return hundred
	}
	fmt.Println("Big number:", addHundred())
	fmt.Println("Big number:", addHundred())
}

// Function that takes an arbitrary number of arguments
func calculateAverage(numbers ...float64) float64 {
	fmt.Println(numbers)
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

// Function that takes an arbitrary number of arguments of any type
// They are called "Variadic Funcitons"
func boringFunction(someBoringStuff ...interface{}) {
	fmt.Println(someBoringStuff...)
}
