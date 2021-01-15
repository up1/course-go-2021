package main

import "fmt"

func main() {
	// Declare an array of five strings that is initialized
	var fruits [5]string
	fruits[0] = "Apple"

	// Iterate over the array
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%d, [%s]\n", i, fruits[i])
	}

	// Iterate over the array with range
	for i, fruit := range fruits {
		fmt.Printf("%d, [%s]\n", i, fruit)
	}
}
