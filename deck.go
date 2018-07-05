package main

import (
	"fmt"
)

// creates new type of deck

type deck []string

// this is a proper deck
func newDeck() deck {

	cards := deck{}

	suits := []string{" of Hearts", " of Dimonds", " of Clubs", " of Spades"}
	cardNo := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "nine",
		"Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, card := range cardNo {

			cards = append(cards, card+suit)

		}
		fmt.Println()
	}

	return cards
}

//reciver Allows the used of a func by any varible of type
//i.e in this case any var of typpe deck will have access to these funcs
//reciver is a var that is passed before the func name annotation
func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	fmt.Println("shuffle", d)
}

func (d deck) deal() {

}
