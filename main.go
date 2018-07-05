package main

import (
	"fmt"
)

func main() {

	cards := newDeck()

	for i := 0; i < len(cards); i++ {

		fmt.Println(cards[i])
	}

}
