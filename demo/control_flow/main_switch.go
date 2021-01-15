package main

import "fmt"

func main() {

	name := "Jon"

	// Switch cases can have multiple values and do not require breaks.
	switch name {
	case "Cersei", "Jaime", "Tyrion":
		fmt.Println("A Lannister always pays their debts.")
	case "Robb", "Sansa", "Bran", "Arya", "Rickon":
		fmt.Println("Winter is coming.")
	default:
		fmt.Println("A man has no name.")
	}

	age := 32

	// A switch with no expression is equivalent to "switch true".
	switch {
	case age < 18:
		fmt.Println("Kids rate is $5.00")
	case age >= 18 && age < 55:
		fmt.Println("Adults rate is $7.00")
	case age >= 55:
		fmt.Println("Seniors rate is $6.00")
	}
}
