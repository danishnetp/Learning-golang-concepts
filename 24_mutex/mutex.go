package main

import (
	"fmt"
	"sync"
)

// post represents shared state that many goroutines may access at once.
//
// views is a shared counter.
// mu protects views so only one goroutine updates it at a time.
type post struct {
	views int
	mu    sync.Mutex // Mutex to protect access to views
}

// inc increments the shared views counter exactly once.
//
// Concurrency notes:
// - Lock/Unlock create a critical section around the increment.
// - Without the mutex, concurrent writes to views would race.
// - wg.Done is deferred so Wait always eventually unblocks.
func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	defer p.mu.Unlock()
	p.views += 1
}

func main() {
	// Initialize shared object with zero views.
	p := post{views: 0}
	var wg sync.WaitGroup

	// Launch 100 goroutines; each goroutine increments the shared counter once.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.inc(&wg)
	}

	// Wait for all workers to finish, then print a deterministic final count.
	wg.Wait()
	fmt.Println("Views:", p.views)
}
