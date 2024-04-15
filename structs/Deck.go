package structs

import (
	"sync"
)

type CardNode struct {
	data *Card
	next *CardNode
}

// Deck represents a concurrency-safe Last-In-First-Out (Stack) container for cards.
type Deck struct {
	lock *sync.Mutex // Mutex to ensure concurrent safety.
	head *CardNode   // Pointer to the top node of the stack.
	Size int         // Current size of the stack.
}

// Push adds a new card onto the top of the deck.
func (d *Deck) Push(data *Card) {
	d.lock.Lock()         // Lock acquired for writing
	defer d.lock.Unlock() // Ensure the lock is released after the function exits

	newNode := &CardNode{ // Create a new node
		data: data,
	}

	newNode.next = d.head // Set the next pointer of the new node to the current head
	d.head = newNode      // Update the head to point to the new node
	d.Size++              // Increment the size of the deck
}

// Pop removes and returns the top card from the deck.
func (d *Deck) Pop() *Card {
	d.lock.Lock()         // Lock acquired for writing
	defer d.lock.Unlock() // Ensure the lock is released after the function exits

	if d.head == nil { // Check if the deck is empty
		return nil
	}

	poppedCard := d.head.data // Retrieve the data from the current head node
	d.head = d.head.next      // Move the head pointer to the next node
	d.Size--                  // Decrement the size of the deck

	return poppedCard // Return the popped card
}

// NewDeck creates a new empty deck.
func NewDeck() *Deck {
	d := &Deck{
		lock: &sync.Mutex{}, // Initialize the mutex for concurrent safety.
	}

	return d
}
func (d *Deck) Iterate() <-chan *Card {
	ch := make(chan *Card)

	go func() {
		defer close(ch) // Ensure the channel is closed when we finish sending items

		d.lock.Lock()         // Lock acquired for reading
		defer d.lock.Unlock() // Ensure the lock is released after the function exits

		current := d.head // Start at the head of the deck

		for current != nil {
			ch <- current.data     // Send the current card to the channel
			current = current.next // Move to the next card
		}
	}()

	return ch
}
