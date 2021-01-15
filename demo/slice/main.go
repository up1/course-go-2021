package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	var s []int = numbers[1:3]
	fmt.Println(s)

	names := make([]string, 2)
	names[0] = "n1"
	names[1] = "n2"
	names[2] = "Error" // Runtime error
	names = append(names, "n3")
	fmt.Println(names)
}
