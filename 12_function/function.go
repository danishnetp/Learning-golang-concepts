package main

import "fmt"

func main() {
	result := add(3, 5)
	fmt.Println("The sum is:", result)
}

// Create a function that takes two integers as input and returns their sum.
func add(a int, b int) int {
	return a + b
}
