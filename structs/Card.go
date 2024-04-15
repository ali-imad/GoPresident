package structs

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Suit uint8

const (
	Jack = iota + 11
	Queen
	King
	Ace
)

const (
	Diamond Suit = iota
	Club
	Heart
	Spade
)

type Card struct {
	suit Suit
	num  uint8
}

func (c Card) String() string {
	var numStr = ""
	var suitStr = ""
	if c.suit == Heart {
		suitStr = "Hearts"
	}
	if c.suit == Club {
		suitStr = "Clubs"
	}
	if c.suit == Spade {
		suitStr = "Spades"
	}
	if c.suit == Diamond {
		suitStr = "Diamonds"
	}
	if c.num == Jack {
		numStr = "Jack"
	} else if c.num == Queen {
		numStr = "Queen"
	} else if c.num == King {
		numStr = "King"
	} else if c.num == Ace {
		numStr = "Ace"
	} else {
		numStr = strconv.Itoa(int(c.num))
	}
	return fmt.Sprintf("%s of %s", numStr, suitStr)
}

// CreateShuffledDeck creates a new deck with cards in a shuffled order.
func CreateShuffledDeck() *Deck {
	cards := createCards() // Create an array of standard cards (A-2, S - D - H - C)
	shuffle(cards)         // Shuffle the array of cards
	deck := NewDeck()      // Create a new empty deck

	for i := 0; i < len(cards); i++ {
		deck.Push(cards[i])
	}

	return deck
}

// createCards creates an array of cards.
func createCards() []*Card {
	var cards []*Card
	index := 0
	for suit := Diamond; suit <= Spade; suit++ {
		for num := 2; num <= Ace; num++ {
			card := &Card{suit: suit, num: uint8(num)}
			cards = append(cards, card)
			index++
		}
	}
	return cards
}

// shuffle shuffles the given array of cards.
func shuffle(cards []*Card) {
	// Use Fisher-Yates shuffle algorithm
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}
