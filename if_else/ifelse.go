package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		println("You are eligible to vote")
	} else if age >= 12 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}

	// We can declare a variable in the if statement, and it will be available in all branches of the if statement.
	if age := 18; age >= 18 {
		fmt.Println("You are eligible to vote, current age: ", age)
	} else if age >= 12 {
		fmt.Println("You are a teenager, current age: ", age)
	} else {
		fmt.Println("You are a child, current age: ", age)
	}

	// Go does not have a ternary operator, but we can use if-else statement to achieve the same result.
}
