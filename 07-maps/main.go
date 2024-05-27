package main

import "fmt"

type User struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	// Maps are like Java's HashTables
	grades := make(map[string]int)
	grades["Harry Potter"] = 23
	grades["John Doe"] = 34
	grades["Frodo Baggins"] = 30
	fmt.Println("Grades:", grades)

	delete(grades, "Frodo Baggins")
	fmt.Println("Grades:", grades)

	userDatabase := make(map[string]*User)

	userOne := User{FirstName: "Michael", LastName: "Jackson", Age: 54}
	userTwo := User{FirstName: "Albus", LastName: "Dumbledore", Age: 70}
	userThree := User{FirstName: "Jon", LastName: "Snow", Age: 34}

	userDatabase["1"] = &userOne
	userDatabase["2"] = &userTwo
	userDatabase["3"] = &userThree

	fmt.Println("User database:", userDatabase)
	userDatabase["1"].FirstName = "Jermaine"
	fmt.Println(userDatabase["1"].FirstName, "and ", userOne.FirstName)

	for key, value := range userDatabase {
		fmt.Println("user id:", key, "First Name:", value.FirstName)
	}

}
