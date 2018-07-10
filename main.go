package main

func main() {

	cards := newDeck()

	hand, remaining := deal(cards, 3)

	//hand.print()

	//fmt.Println()

	//remaining.print()

	hand.saveDeck("myHand")
	remaining.saveDeck("restOfDeck")

	newHand := openFile("myHand")

	newHand.print()
}
