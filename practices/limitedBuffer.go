package main

import (
	"fmt"
	"sync"
)

// RingBuffer represents a fixed-size, thread-safe FIFO buffer
type RingBuffer struct {
	mu       sync.Mutex
	messages []string
	size     int
	head     int // Points to the oldest message
	tail     int // Points to where the next new message will go
	count    int // Tracks current number of items in the buffer
}

// NewRingBuffer initializes a buffer with a strict maximum size
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		messages: make([]string, size),
		size:     size,
	}
}

// Push adds a new message. If full, it overwrites the oldest message automatically.
func (b *RingBuffer) Push(msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Insert the message at the current tail position
	b.messages[b.tail] = msg

	// Advance the tail pointer circularly
	b.tail = (b.tail + 1) % b.size

	if b.count < b.size {
		// Buffer is not full yet, just increment total count
		b.count++
	} else {
		// Buffer is full! The oldest message was just overwritten.
		// Move the head forward to the next oldest message.
		b.head = (b.head + 1) % b.size
	}
}

// GetAll retrieves all active messages from oldest to newest without clearing the buffer
func (b *RingBuffer) GetAll() []string {
	b.mu.Lock()
	defer b.mu.Unlock()

	result := make([]string, b.count)
	current := b.head

	for i := 0; i < b.count; i++ {
		result[i] = b.messages[current]
		current = (current + 1) % b.size
	}

	return result
}

func main() {
	// 1. Create a buffer limited to exactly 3 messages
	buf := NewRingBuffer(3)

	// 2. Add 3 messages (Buffer becomes full)
	buf.Push("Msg 1")
	buf.Push("Msg 2")
	buf.Push("Msg 3")
	fmt.Println("Initial buffer:", buf.GetAll())
	// Output: [Msg 1 Msg 2 Msg 3]

	// 3. Add a 4th message. "Msg 1" (oldest) will be dropped.
	buf.Push("Msg 4")
	fmt.Println("After 4th push:", buf.GetAll())
	// Output: [Msg 2 Msg 3 Msg 4]

	// 4. Add a 5th message. "Msg 2" (oldest) will be dropped.
	buf.Push("Msg 5")
	fmt.Println("After 5th push:", buf.GetAll())
	// Output: [Msg 3 Msg 4 Msg 5]
}
