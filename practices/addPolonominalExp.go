package main

import (
	"fmt"
)

// Expression represents an algebraic expression of the form: ax + by + c
type Expression struct {
	X        int
	Y        int
	Constant int
}

// Add adds two expressions together
func (e1 Expression) Add(e2 Expression) Expression {
	return Expression{
		X:        e1.X + e2.X,
		Y:        e1.Y + e2.Y,
		Constant: e1.Constant + e2.Constant,
	}
}

// Subtract subtracts the second expression from the first
func (e1 Expression) Subtract(e2 Expression) Expression {
	return Expression{
		X:        e1.X - e2.X,
		Y:        e1.Y - e2.Y,
		Constant: e1.Constant - e2.Constant,
	}
}

// String formats the expression cleanly for printing
func (e Expression) String() string {
	res := fmt.Sprintf("%dx", e.X)

	// Handle Y sign formatting
	if e.Y >= 0 {
		res += fmt.Sprintf("+%dy", e.Y)
	} else {
		res += fmt.Sprintf("%dy", e.Y)
	}

	// Handle Constant sign formatting
	if e.Constant >= 0 {
		res += fmt.Sprintf("+%d", e.Constant)
	} else {
		res += fmt.Sprintf("%d", e.Constant)
	}

	return res
}

func main() {
	// Define the input expressions
	expr1 := Expression{X: 5, Y: 27, Constant: 10}
	expr2 := Expression{X: 2, Y: -1, Constant: 2}

	// Perform operations
	additionResult := expr1.Add(expr2)
	subtractionResult := expr1.Subtract(expr2)

	// Print results
	fmt.Printf("Input:\n%s and %s\n\n", expr1, expr2)
	fmt.Printf("Addition: %s\n", additionResult)
	fmt.Printf("Subtraction: %s\n", subtractionResult)
}
