package main

import (
	"fmt"
	"sync"
)

// Race condition background:
// A race happens when multiple goroutines read/write shared state at the same
// time without synchronization, and at least one operation is a write.
//
// In this file, many goroutines increment the same counter concurrently.
// Because the increment is not protected, some updates can be lost.
// Expected value is 100, but actual output can be lower or nondeterministic.
//
// This example intentionally demonstrates the unsafe version first.
// A mutex-based safe version would lock before incrementing and unlock after.

type post struct {
	views int
}

// inc increments the shared views counter once.
// This function is intentionally NOT synchronized to demonstrate a race.
func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views += 1
}

func main() {
	// Shared post value that every goroutine mutates.
	p := post{views: 0}

	// WaitGroup ensures main waits until all goroutines finish.
	var wg sync.WaitGroup

	// Launch 100 goroutines, each incrementing p.views once.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.inc(&wg)
	}

	// Wait for all increments to complete.
	wg.Wait()

	// Without synchronization, this may not always print 100.
	fmt.Println("Views:", p.views)
}
