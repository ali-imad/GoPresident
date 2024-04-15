package test

import (
	"president-cli/structs"
	"testing"
)

func TestCardString(t *testing.T) {
	c := &structs.Card{Suit: structs.Spade, Num: structs.Ace}
	ex := "Ace of Spades"

	if c.String() != ex {
		t.Errorf("Expected %s, got %s", ex, c.String())
	}

	c2 := &structs.Card{Suit: structs.Heart, Num: structs.Jack}
	ex2 := "Jack of Hearts"

	if c2.String() != ex2 {
		t.Errorf("Expected %s, got %s", ex2, c2.String())
	}

	c3 := &structs.Card{Suit: structs.Club, Num: 9}
	ex3 := "9 of Clubs"

	if c3.String() != ex3 {
		t.Errorf("Expected %s, got %s", ex3, c3.String())
	}

	c4 := &structs.Card{Suit: structs.Diamond, Num: structs.Queen}
	ex4 := "Queen of Diamonds"

	if c4.String() != ex4 {
		t.Errorf("Expected %s, got %s", ex4, c3.String())
	}

	c5 := &structs.Card{Suit: structs.Club, Num: structs.King}
	ex5 := "King of Clubs"

	if c5.String() != ex5 {
		t.Errorf("Expected %s, got %s", ex5, c3.String())
	}
}
