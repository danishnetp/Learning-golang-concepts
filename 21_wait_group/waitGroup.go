package main

import (
	"fmt"
	"sync"
)

// waitGroupExample demonstrates how sync.WaitGroup coordinates multiple goroutines.
//
// Flow summary:
// 1) Add(n) declares how many goroutines must finish.
// 2) Each goroutine calls Done() exactly once when its work ends.
// 3) Wait() blocks until the internal counter returns to zero.
//
// This pattern prevents main logic from exiting before background work completes.
func waitGroupExample() {
	// WaitGroup is a counter-based synchronization primitive.
	var wg sync.WaitGroup

	// Add must be called before launching worker goroutines.
	// Here we plan to run exactly 5 tasks.
	wg.Add(5)

	// Start 5 concurrent workers.
	for i := 1; i <= 5; i++ {
		// Pass i as a parameter to avoid closure capture issues with loop variables.
		go func(id int) {
			// Done decrements the WaitGroup counter by 1.
			// defer ensures it runs even if the function returns early.
			defer wg.Done()

			// Simulated unit of work.
			fmt.Printf("Doing the task %d\n", id)
		}(i)
	}

	// Wait blocks until all 5 Done calls have been executed.
	wg.Wait()
}

func main() {
	// main launches the synchronized goroutine demo and then exits cleanly.
	fmt.Println("Starting wait group example...")
	waitGroupExample()
	fmt.Println("All tasks completed.")
}
