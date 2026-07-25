package main

import "fmt"

// main compares passing by value and passing by reference using pointers.
//
// callByValue receives a copy, so the original variable is unchanged.
// callByReference receives an address, so it can modify the original value.
func main() {
	x := 5
	fmt.Println("Before callByValue, x:", x) // Output: Before callByValue, x: 5
	callByValue(x)
	fmt.Println("After callByValue, x:", x) // Output: After callByValue, x: 5

	fmt.Println("Before callByReference, x:", x) // Output: Before callByReference, x: 5
	callByReference(&x)
	fmt.Println("After callByReference, x:", x) // Output: After callByReference, x: 15
}

// callByValue adds 10 to a local copy of x.
//
// Because x is passed by value, changes inside this function do not affect
// the caller's variable.
func callByValue(x int) {
	x = x + 10
	fmt.Println("Inside callByValue, x:", x)
}

// callByReference adds 10 to the value stored at the provided address.
//
// Because x is a pointer, this function updates the caller's original value.
func callByReference(x *int) {
	*x = *x + 10
	fmt.Println("Inside callByReference, x:", *x)
}
