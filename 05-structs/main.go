package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Country   string
}

type Book struct {
	Title  string
	Author string
}

func main() {
	personOne := Person{
		FirstName: "Frodo",
		LastName:  "Baggins",
		Country:   "Shire",
		// we didn't set the Age, it will default to zero
	}

	fmt.Println("PersonOne", personOne)
	fmt.Println("PersonOne full name", personOne.getFullName())

	bookOne := Book{
		Title:  "Lord of the Rings",
		Author: "Tolkien",
	}
	fmt.Println("bookOne title:", bookOne.Title)
	// changeTitle(&bookOne, "Harry Potter")
	// fmt.Println("bookOne title:", bookOne.Title)
	bookOne.changeTitleAsAReceiver("Harry Potter")
	fmt.Println("bookOne title", bookOne.Title)
}

// The parantheses at the beginning which takes a pointer to a Person is called scoping
// It's basically scopes the function to the struct Person. This way the function is like a
// member function of the struct --> This is called a Receiver or a Method
func (person *Person) getFullName() string {
	return person.FirstName + " " + person.LastName
}

func changeTitle(book *Book, newTitle string) {
	book.Title = newTitle
}

func (book *Book) changeTitleAsAReceiver(newTitle string) {
	book.Title = newTitle
}
