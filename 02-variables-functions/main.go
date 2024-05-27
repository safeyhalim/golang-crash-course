package main

import "fmt"

var globalVariable = "This is a global variable" // with var but without :=

func main() {
	var awesomeString string // declaration of a variable
	var reallyCoolInteger int

	awesomeString = "Hello awesome world!" // assignment of of awesomeString
	reallyCoolInteger = 23
	coolFloat := 34.3 // declaration and assignment in the same step (no need for a var). Needs a walrus ":="

	fmt.Println("awesomeString is:", awesomeString)
	fmt.Println("reallyCoolInteger is:", reallyCoolInteger)
	fmt.Println("coolFloat is:", coolFloat)
	fmt.Println("globalVariable is:", globalVariable)
	fmt.Println("firstName is :", getFirstName())

	bob, carol := getABunchOfNames()
	fmt.Println("Name one:", bob)
	fmt.Println("Name two:", carol)

	incrementedNumber := incrementByOne(11)
	fmt.Println("incrementedNumber:", incrementedNumber)

	ourNewFloat := addTwoFloats(13.3, 45.3)
	fmt.Println("ourNewFloat is:", ourNewFloat)
}

func getFirstName() string {
	return "Safey"
}

func getABunchOfNames() (string, string) { // Function that returns more than one value
	return "Bob", "Carol"
}

func incrementByOne(number int) int {
	return number + 1
}

func addTwoFloats(numberOne, numberTwo float64) float64 { // When the two parameters are of the same type, the type should only be declared once
	return numberOne + numberTwo
}

func SubtractTwoIntegers(numberOne, numberTwo int) int { // The function starts with a capital letter, this makes it public: it can be used in other files
	return numberOne - numberTwo
}
