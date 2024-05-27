package main

import (
	"fmt"
	"time"
)

// Concurrency in Go
// Goroutines are lightweight threads: You can execute 100s of 1000s of them

func doSomething(thingToDo string) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Step:", i, "I am ", thingToDo)
		time.Sleep(time.Millisecond * 500)
	}
}

// The main function itself is an implicit Goroutine (thread #1)
func main() {
	go doSomething("eating breakfast")           // creates another Goroutine (thread #2)
	go doSomething("watching Netflix")           // thread #3
	go doSomething("starring out of the window") // thread #4

	fmt.Scanln()
}
