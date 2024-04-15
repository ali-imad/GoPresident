package structs

import (
	"reflect"
	"sync"
	"testing"
)

func TestDeck_Iterate(t *testing.T) {
	// Create a new deck
	deck := NewDeck()
	// Populate the deck with 52 cards in known order
	for suit := Diamond; suit <= Spade; suit++ {
		for num := MinNum; num <= Ace; num++ {
			card := &Card{Suit: suit, Num: num}
			deck.Push(card)
		}
	}
	// Iterate through the deck and comparing its .Pop() to the card building sequence in reverse order
	i := 0
	for card := range deck.Iterate() {
		expectedCard := &Card{Suit: Spade - Suit(i/TotalNums), Num: Ace - i%TotalNums}
		if !reflect.DeepEqual(card, expectedCard) {
			t.Errorf("Expected %v, got %v\n", expectedCard, card)
		}
		i++
	}
}

func TestDeck_CreateShuffledDeck(t *testing.T) {
	shuffledDeck := NewDeck()
	if !(shuffledDeck.Pop() == nil) {
		t.Errorf("Expected empty deck, got non-empty deck\n")
	}
	shuffledDeck = CreateShuffledDeck() // Create a shuffled deck
	unshuffledDeck := NewDeck()         // Create an unshuffled deck

	// Fill the unshuffled deck with cards in a known order
	for suit := Diamond; suit <= Spade; suit++ {
		for num := MinNum; num <= Ace; num++ {
			card := &Card{Suit: suit, Num: num}
			unshuffledDeck.Push(card)
		}
	}

	// Check the size of the deck
	if shuffledDeck.Size != 52 {
		t.Errorf("Expected deck size to be 52, got %d\n", shuffledDeck.Size)
	}

	// Check if the deck is shuffled by comparing it to the unshuffled deck
	var isNotShuffled = true
	for i := 0; i < 52; i++ {
		shuffledCard := shuffledDeck.Pop()
		unshuffledCard := unshuffledDeck.Pop()

		if !reflect.DeepEqual(shuffledCard, unshuffledCard) {
			isNotShuffled = false
			break
		}
	}
	if isNotShuffled {
		t.Errorf("Expected the deck to be shuffled, got same cards\n")
	}
}

func TestDeck_ConcurrentPushPop(t *testing.T) {
	const numPushers = 1000 // Number of goroutines pushing to the deck
	const numPoppers = 1000 // Number of goroutines popping from the deck
	const totalOps = 10000  // Total number of push and pop operations
	const initSize = 2000

	deck := NewDeck() // Create a new concurrent-safe deck

	// Initialize the deck to create a buffer
	for i := 0; i < initSize; i++ {
		deck.Push(CreateRandomCard())
	}

	// Wait group to synchronize all goroutines
	var wg sync.WaitGroup
	wg.Add(numPushers + numPoppers)

	// Goroutines for pushing to the deck
	for i := 0; i < numPushers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < totalOps/numPushers; j++ {
				deck.Push(&Card{}) // Push a card to the deck
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
