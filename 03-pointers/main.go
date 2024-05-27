package main

import "fmt"

func main() {
	var myFavouriteFruit string

	myFavouriteFruit = "apple"
	fmt.Println("My favourite fruit is:", myFavouriteFruit)
	changeFavouriteFruit(&myFavouriteFruit, "peach") // passing the address to the function
	fmt.Println("My favourite fruit is:", myFavouriteFruit)
}

func changeFavouriteFruit(fruitPointer *string, newFavouriteFruit string) { // pointer to a string
	fmt.Println("Address of the fruitPoint is:", fruitPointer, "the value is:", *fruitPointer) // printing the address of the pointer then the value it points to
	*fruitPointer = newFavouriteFruit                                                          // changing the value of the pointer
}
