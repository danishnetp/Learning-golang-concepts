package main

import (
	"fmt"
)

// printSlice accepts a slice of any element type and prints each item.
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// printSliceWithMultipleTypes demonstrates multiple type parameters.
// T must be comparable, and U is constrained to string.
func printSliceWithMultipleTypes[T comparable, U string](items []T, named U) {
	for _, item := range items {
		fmt.Println(item, named)
	}
}

// printSliceWithBounds uses a stricter constraint (comparable) for T.
// This allows only types that can be compared with == and !=.
func printSliceWithBounds[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// printSliceInterface is equivalent to printSlice in this example.
// Using interface{} would work, but generics provide type safety and avoid
// unnecessary type assertions at call sites.
func printSliceInterface[T interface{}](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// stack is a generic struct that stores elements of one type.
type stack[T any] struct {
	elements []T
}

func main() {
	// 1) Generic function used with int and string slices.
	ints := []int{1, 2, 3, 4, 5}
	strings := []string{"apple", "banana", "cherry"}

	printSlice(ints)    // Output: 1 2 3 4 5
	printSlice(strings) // Output: apple banana cherry

	// 2) Same behavior through a generic function constrained by interface{}.
	printSliceInterface(ints)    // Output: 1 2 3 4 5
	printSliceInterface(strings) // Output: apple banana cherry

	// 3) Constraint-based generic function examples.
	printSliceWithBounds([]bool{true, false, true}) // Output: true false true
	printSliceWithBounds([]float64{1.1, 2.2, 3.3})  // Output: 1.1 2.2 3.3
	printSliceWithBounds(strings)                   // Output: apple banana cherry

	// 4) Generic struct instantiated with int.
	myStack := stack[int]{elements: []int{10, 20, 30}}
	fmt.Println("Stack elements:", myStack.elements) // Output: Stack elements: [10 20 30]

	// 5) Generic struct instantiated with string.
	myStringStack := stack[string]{elements: []string{"x", "y", "z"}}
	fmt.Println("String Stack elements:", myStringStack.elements) // Output: String Stack elements: [x y z]

	// 6) Function with multiple type parameters.
	printSliceWithMultipleTypes(ints, "numbers")
}
