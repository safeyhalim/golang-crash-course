package main

import "fmt"

func main() {
	var number int
	number = 15 // I will be happy again, as happy as this world allows me to be

	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else if number%3 == 0 {
		fmt.Println(number, "is a multiple of 3")
	} else {
		fmt.Println(number, "is odd")
	}

	// Switch Statements
	favouriteFruit := "apple"
	switch favouriteFruit {
	case "apple":
		fmt.Println("Yum apples are really nice")
		break
	case "pear":
		fmt.Println("Hmmm pears are nice too")
		break
	case "banana":
		fmt.Println("You can slip on them, but they are nice")
		break
	default:
		fmt.Println("I don't think this is a fruit")
	}

	// For loops (Go doesn't have a concept of a while statement. Only several forms of for loops)
	i := 1
	for i <= 10 {
		fmt.Println(i, "doubled is:", i+i)
		i++
	}

	for x := 1; x <= 20; x++ {
		if x%2 == 0 {
			fmt.Println(x, "is even")
		}
		// fmt.Println(x, "squared is:", x*x)
	}

}
