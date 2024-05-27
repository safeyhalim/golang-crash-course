package main

import (
	"fmt"
	"time"
)

// Channels is a way for Goroutines to communicate with each other

func giveMeNumber(intChannel chan int) { // takes a channel of integers
	defer close(intChannel) // Close the channel at the end of the function (tke keyword: defer)
	for i := 0; i <= 10; i++ {
		intChannel <- i // throws i into the intChannel
	}
}

func doubleNumber(intChannel chan int) { // takes a channel of integers
	for number := range intChannel { // throws the values of the channel in the variable number until the channel is closed
		fmt.Println(number, "doubled is:", number*2)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	intChannel := make(chan int) // Create a channel of integers
	go giveMeNumber(intChannel)  // thread #2
	doubleNumber(intChannel)     // in thread #1
}
