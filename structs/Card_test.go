package structs

import (
	"testing"
)

func TestCard_String(t *testing.T) {
	c := &Card{Suit: Spade, Num: Ace}
	ex := "A.s"

	if c.String() != ex {
		t.Errorf("Expected %s, got %s", ex, c.String())
	}

	c2 := &Card{Suit: Heart, Num: Jack}
	ex2 := "J.h"

	if c2.String() != ex2 {
		t.Errorf("Expected %s, got %s", ex2, c2.String())
	}

	c3 := &Card{Suit: Club, Num: 9}
	ex3 := "9.c"

	if c3.String() != ex3 {
		t.Errorf("Expected %s, got %s", ex3, c3.String())
	}

	c4 := &Card{Suit: Diamond, Num: Queen}
	ex4 := "Q.d"

	if c4.String() != ex4 {
		t.Errorf("Expected %s, got %s", ex4, c4.String())
	}

	c5 := &Card{Suit: Club, Num: King}
	ex5 := "K.c"

	if c5.String() != ex5 {
		t.Errorf("Expected %s, got %s", ex5, c5.String())
	}
}

func TestCard_LongString(t *testing.T) {
	c := &Card{Suit: Spade, Num: Ace}
	ex := "Ace of Spades"

	if c.LongString() != ex {
		t.Errorf("Expected %s, got %s", ex, c.LongString())
	}

	c2 := &Card{Suit: Heart, Num: Jack}
	ex2 := "Jack of Hearts"

	if c2.LongString() != ex2 {
		t.Errorf("Expected %s, got %s", ex2, c2.LongString())
	}

	c3 := &Card{Suit: Club, Num: 9}
	ex3 := "9 of Clubs"

	if c3.LongString() != ex3 {
		t.Errorf("Expected %s, got %s", ex3, c3.LongString())
	}

	c4 := &Card{Suit: Diamond, Num: Queen}
	ex4 := "Queen of Diamonds"

	if c4.LongString() != ex4 {
		t.Errorf("Expected %s, got %s", ex4, c3.LongString())
	}

	c5 := &Card{Suit: Club, Num: King}
	ex5 := "King of Clubs"

	if c5.LongString() != ex5 {
		t.Errorf("Expected %s, got %s", ex5, c3.LongString())
	}
}
