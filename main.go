package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remaining := deal(cards, 3)

	//hand.print()

	//fmt.Println()

	//remaining.print()

	st := hand.saveDeck()
	cmd := remaining.saveDeck()

	fmt.Println(st)
	fmt.Println()
	fmt.Println(cmd)
}
