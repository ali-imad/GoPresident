package view

import (
	"fmt"
	"president-cli/structs"
)

func PrintDeck(d *structs.Deck) {
	i := 1

	for card := range d.Iterate() {
		fmt.Printf("%d:   %s", i, card)
		if i%5 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" \t//\t")
		}
		i++
	}
}
