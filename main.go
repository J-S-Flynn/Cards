package main

import (
	"fmt"
)

func main() {

	cards := deck{"Ace of Spades", newCard()}

	cards = append(cards, "Ace of Clubs")
	cards = append(cards, "Ace of Dimonds")

	//difrent ways of defining a for loop in go
	//type one
	for i, card := range cards {
		fmt.Println(i, card)
	}

	//type two
	for i := 0; i < len(cards); i++ {

		fmt.Println(cards[i])
	}
}

func newCard() string {

	return "Ace of Dimonds"
}
