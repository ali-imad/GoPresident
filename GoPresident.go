package main

import (
	"fmt"
	"president-cli/structs"
)

func main() {
	fmt.Printf("Welcome to GoPresident!\n")
	x := 20
	for x > 0 {
		card := structs.CreateRandomCard()
		fmt.Println(*card)
		x--
	}
	d := structs.CreateShuffledDeck()
	i := 0

	//dumpDeck
	for d.Size > 0 {
		fmt.Printf("%d:\t\t%s\n", i, d.Pop())
		i++
	}
}
