package main

import (
	"fmt"
)

// main demonstrates calling a variadic function with direct values and with a slice.
func main() {
	// Call sum with individual integer arguments.
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", result)

	// Call sum with a slice by expanding it using the ... operator.
	numbers := []int{10, 20, 30}
	result2 := sum(numbers...)
	fmt.Println("The sum of the slice is:", result2)
}

// sum returns the total of any number of integer arguments.
//
// The parameter nums ...int is variadic, so callers can pass zero or more ints.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
