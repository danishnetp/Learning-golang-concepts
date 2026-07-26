package main

import (
	"fmt"
)

// task simulates a unit of work executed by a goroutine.
func task(id int) {
	fmt.Printf("Doing the task %d\n", id)
}

func main() {
	// A goroutine runs concurrently with the main goroutine.
	fmt.Println("Starting goroutines...")
	// Launch 5 goroutines to perform tasks concurrently.
	for i := 1; i <= 5; i++ {
		// Start task(i) in a separate goroutine.
		go task(i)

		// Start an inline anonymous function as another goroutine.
		// Note: this closure captures the loop variable i by reference.
		go func() {
			fmt.Println("Doing an inline task", i)
		}()
	}

	// Block main so goroutines get time to run before program exit.
	fmt.Scanln()
}
