package main

import (
	"fmt"
	"president-cli/structs"
)

func main() {
	fmt.Printf("Welcome to GoPresident!\n")
	d := structs.CreateShuffledDeck()
	i := 0
	for d.Size > 0 {
		fmt.Printf("%d:\t\t%s\n", i, d.Pop())
		i++
	}
}
