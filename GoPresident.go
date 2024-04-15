package main

import (
	"fmt"
	"president-cli/structs"
	"president-cli/view"
)

func main() {
	fmt.Printf("Welcome to GoPresident!\n")
	//x := 20
	//for x > 0 {
	//	card := structs.CreateRandomCard()
	//	fmt.Println(*card)
	//	x--
	//}
	d := structs.CreateShuffledDeck()
	view.PrintDeck(d)
}
