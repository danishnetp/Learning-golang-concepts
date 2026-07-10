package main

import "fmt"

// Constants are immutable values which are known at compile time and do not change for the life of the program.
// Constant variable could be outside the function and inside the function.
// Constant variable could be of any type like int, float, string, boolean etc.
const hello string = "Hello"

func main() {
	// Constant declare with 'const' keyword
	const pi float64 = 3.144
	fmt.Println(pi)
	const age = 30
	fmt.Println(age)
	fmt.Println(hello)

	// Constant with multiple values
	const (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println(a, b, c)

}
