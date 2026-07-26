package main

import (
	"fmt"
	"math/rand"
	"time"
)

// processNum receives integers from numChain and processes them one by one.
//
// Key ideas:
// - `for num := range numChain` keeps receiving until the channel is closed.
// - Each value is handled sequentially in this goroutine.
// - This function acts like a worker that consumes jobs from a shared stream.
//
// In this demo we do not close numChain, so this goroutine can keep waiting for
// more values. The program still exits when main returns.
func processNum(numChain chan int) {
	for num := range numChain {
		fmt.Println("Processing : ", num)
		time.Sleep(time.Second)
	}
}

// sum computes num1 + num2 and sends the result through the result channel.
// This shows how channels can be used to return values from goroutines safely.
// The caller waits on the same channel to receive the computed result.
func sum(result chan int, num1 int, num2 int) {
	result <- num1 + num2 // Send the sum to the result channel.
}

// task simulates background work and signals completion on the done channel.
// This is a common pattern when the caller only needs a completion notification.
func task(done chan bool) {
	defer func() {
		done <- true // Signal that the task is complete.
	}()

	fmt.Println("Processing..")
}

// emailSender reads email addresses from emailChannel until it is closed.
// Directional channel types make intent explicit:
// - <-chan string means receive-only
// - chan<- bool means send-only
// This helps document the API and prevents accidental misuse in the function.
func emailSender(emailChannel <-chan string, done chan<- bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChannel {
		fmt.Println("Sending email to:", email)
		time.Sleep(time.Second) // Simulate time taken to send an email.
	}
}

func main() {
	// Example 1: unbuffered channel used to stream work to a goroutine.
	// numChain is an unbuffered channel of int.
	// Unbuffered send/receive operations synchronize sender and receiver.
	numChain := make(chan int)

	// Launch a goroutine to process numbers from the channel.
	go processNum(numChain)

	// Send five random numbers into numChain.
	// Because numChain is unbuffered, each send waits until processNum receives.
	for i := 0; i < 5; i++ {
		numChain <- rand.Intn(100) // Send a random number to the channel.
	}

	// Example 2: one goroutine computes a value, another receives it.
	// result is another unbuffered channel used for one-shot response.
	result := make(chan int)

	// Launch a goroutine to compute the sum of two numbers.
	go sum(result, 10, 20)

	// Block until a value is received from result.
	res := <-result // Receive the result from the sum channel.
	fmt.Println("Sum is : ", res)

	// Example 3: a done channel signals that some work has completed.
	done := make(chan bool)
	// Launch a goroutine to perform a task and signal completion.
	go task(done)
	<-done // Wait for the task to complete.

	// Example 4: buffered channels allow some sends without an immediate receiver.
	// Buffered channels can hold a limited number of values without blocking.
	emailChannel := make(chan string, 100) // Buffer size of 100

	// Send two email addresses into the buffered channel.
	emailChannel <- "alice@example.com"
	emailChannel <- "bob@example.com"

	// These receives pull values back out of the buffer immediately.
	fmt.Println(<-emailChannel)
	fmt.Println(<-emailChannel)

	// Start a receiver that will consume later values until the channel closes.
	go emailSender(emailChannel, done)

	// Send five more emails for the receiver goroutine to process.
	for i := 0; i < 5; i++ {
		emailChannel <- fmt.Sprintf("user%d@example.com", i)
	}

	fmt.Println("done sending.")

	close(emailChannel) // Close the channel to signal no more emails will be sent.
	<-done              // Wait until emailSender drains the channel and exits.

	// Example 5: select waits on multiple channels and handles whichever is ready.
	// Multiple channels can be used to coordinate different types of data or tasks.
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	// Use select to wait on multiple channel operations.
	// The receive order is nondeterministic and depends on which goroutine sends first.
	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num)
		case msg := <-chan2:
			fmt.Println("Received from chan2:", msg)
		}
	}
}
