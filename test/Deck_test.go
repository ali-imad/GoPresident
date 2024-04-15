package test

import (
	"president-cli/structs"
	"sync"
	"testing"
)

func TestConcurrentPushPop(t *testing.T) {
	const numPushers = 1000 // Number of goroutines pushing to the deck
	const numPoppers = 1000 // Number of goroutines popping from the deck
	const totalOps = 10000  // Total number of push and pop operations
	const initSize = 2000

	deck := structs.NewDeck() // Create a new concurrent-safe deck

	// Initialize the deck to create a buffer
	for i := 0; i < initSize; i++ {
		deck.Push(structs.CreateRandomCard())
	}

	// Wait group to synchronize all goroutines
	var wg sync.WaitGroup
	wg.Add(numPushers + numPoppers)

	// Goroutines for pushing to the deck
	for i := 0; i < numPushers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < totalOps/numPushers; j++ {
				deck.Push(&structs.Card{}) // Push a card to the deck
			}
		}()
	}

	// Goroutines for popping from the deck
	for i := 0; i < numPoppers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < totalOps/numPoppers; j++ {
				deck.Pop() // Pop a card from the deck
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	// Check the final size of the deck
	if deck.Size != initSize {
		t.Errorf("Expected deck size to be %d, got %d", initSize, deck.Size)
	}
}
