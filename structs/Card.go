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
	MinNum = 2
)

const (
	Diamond Suit = iota
	Club
	Heart
	Spade
)

type Card struct {
	Suit Suit
	Num  int
}

func (c Card) String() string {
	var numStr = ""
	var suitStr = ""
	if c.Suit == Heart {
		suitStr = "Hearts"
	}
	if c.Suit == Club {
		suitStr = "Clubs"
	}
	if c.Suit == Spade {
		suitStr = "Spades"
	}
	if c.Suit == Diamond {
		suitStr = "Diamonds"
	}
	if c.Num == Jack {
		numStr = "Jack"
	} else if c.Num == Queen {
		numStr = "Queen"
	} else if c.Num == King {
		numStr = "King"
	} else if c.Num == Ace {
		numStr = "Ace"
	} else {
		numStr = strconv.Itoa(int(c.Num))
	}
	return fmt.Sprintf("%s of %s", numStr, suitStr)
}

// CreateShuffledDeck creates a new deck with cards in a shuffled order.
func CreateShuffledDeck() *Deck {
	cards := createCards() // Create an array of standard cards (2 - A, S - D - H - C)
	shuffle(cards)         // Shuffle the array of cards
	deck := NewDeck()      // Create a new empty deck

	for i := 0; i < len(cards); i++ {
		deck.Push(cards[i])
	}

	return deck
}

// CreateRandomCard is used for testing
func CreateRandomCard() *Card {
	cardSuit := rand.Intn(int(Spade) + 1)
	cardNum := rand.Intn(int(Ace+1-MinNum)) + MinNum
	return &Card{Suit(cardSuit), cardNum}
}

// createCards creates an array of cards.
func createCards() []*Card {
	var cards []*Card
	index := 0
	for suit := Diamond; suit <= Spade; suit++ {
		for num := MinNum; num <= Ace; num++ {
			card := &Card{Suit: suit, Num: num}
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
