package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	var s1 []int = numbers[1:3]
	var s2 []int = numbers[2:4]

	fmt.Println(numbers)
	fmt.Println(s1)
	fmt.Println(s2)

	// Change data
	s2[0] = 333

	fmt.Println("numbers :", numbers)
	fmt.Println("s1 :", s1)
	fmt.Println("s2 :", s2)
}
