package main

import "fmt"

func main() {

	cards := newDeck()

	cards.shuffle()

	hand, remaining := deal(cards, 5)

	hand.print()
	fmt.Println()
	remaining.print()

}
