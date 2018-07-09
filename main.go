package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remaining := deal(cards, 3)

	hand.print()
	fmt.Println()
	remaining.print()
}
