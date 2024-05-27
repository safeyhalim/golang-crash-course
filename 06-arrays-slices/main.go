package main

import (
	"fmt"
	"sort"
)

func main() {
	// Array of 4 elements
	var sweets = [4]string{
		"milk chocolate",
		"dark chocolate",
		"marshmellow",
		"toffee apple",
	}

	fmt.Println("sweet at index 0:", sweets[0])
	sweets[1] = "caramel"
	fmt.Println("sweet at index 1:", sweets[1])
	fmt.Println("length of sweets:", len(sweets))
	fmt.Println("sweets array:", sweets)

	// Slice --> This is like a list
	var scores []int
	scores = append(scores, 10)
	scores = append(scores, 11)
	scores = append(scores, 12)
	scores = append(scores, 13)
	fmt.Println("scores slice:", scores)

	newScores := []int{4, 5, 3, 6, 10} // Another way of adding elements in a slice
	newScores = append(newScores, 13)
	fmt.Println("newsScores slice", newScores)

	// Sortings a slice
	sort.Ints(scores)
	fmt.Println("scores slice sorted", scores)

	// Looping through a slice
	for i := 0; i < len(scores); i++ {
		fmt.Println("index:", i, "value:", scores[i])
	}

	// Better way of looping
	for index, score := range scores {
		fmt.Println("index:", index, "score:", score)
	}

	// If we don't need the index
	for _, score := range scores {
		fmt.Println("score:", score)
	}

	fmt.Println("The average score is:", averageScores(scores))
}

func averageScores(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total / len(scores))
}
