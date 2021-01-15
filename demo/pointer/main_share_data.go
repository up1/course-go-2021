package main

func main() {

	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

func increment(inc *int) {
	// Increment the "value of" count that the "pointer points to".
	*inc++

	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
