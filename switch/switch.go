package main

import (
	"time"
)

func main() {
	i := 5
	switch i {
	case 1:
		println("One")
		// no break statement needed, execution will not fall through to the next case
	case 2:
		println("Two")
	case 3:
		println("Three")
	default:
		println("Default")
	}

	// switch with day of the week
	switch time.Now().Weekday() {
	case time.Monday:
		println("Start of the work week")
	case time.Friday:
		println("End of the work week")
	case time.Saturday, time.Sunday:
		println("Weekend")
	default:
		println("Invalid day")
	}

	// multiple cases can be combined
	switch i {
	case 1, 2, 3:
		println("One, Two or Three")
	default:
		println("Default")
	}

	// switch can also be used without an expression, in which case it acts like an if-else chain
	switch {
	case i < 0:
		println("Negative")
	case i == 0:
		println("Zero")
	case i > 0:
		println("Positive")
	}

	// switch can also be used with type assertions
	var x interface{} = 5
	switch x.(type) {
	case int:
		println("x is an int")
	case string:
		println("x is a string")
	default:
		println("x is of unknown type")
	}

	// swithch with function
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am an int", i.(int))
		case string:
			println("I am a string", i.(string))
		default:
			println("I am of unknown type", i)
		}
	}

	whoAmI("Hello")

}
