package main

import (
	"fmt"
)

// main demonstrates two closure patterns:
// 1) Capturing an outer variable.
// 2) Preserving private state across function calls.
func main() {
	// This closure captures x from the outer scope.
	x := 10
	closure := func(y int) int {
		return x + y
	}
	fmt.Println(closure(5)) // Output: 15

	// counter returns a closure that keeps its own count value.
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment()) // Output: 2
}

// counter returns a closure that increments and returns an internal count.
//
// The count variable is captured by the returned function, so its value
// persists between calls.
func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}
