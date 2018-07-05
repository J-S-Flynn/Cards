package main

import (
	"fmt"
)

func main() {

	cards := []string{"Ace of Spades", newCard()}

	cards = append(cards, "Ace of Clubs")
	cards = append(cards, "Ace of Dimonds")

	for i := 0; i < len(cards); i++ {

		fmt.Println(cards[i])
	}
}

func newCard() string {

	return "Ace of Dimonds"
}
